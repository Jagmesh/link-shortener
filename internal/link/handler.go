package link

import (
	"link-shortener/config"
	"link-shortener/internal/auth"
	"link-shortener/internal/stat"
	"link-shortener/model"
	apperror "link-shortener/pkg/app-error"
	"link-shortener/pkg/jwt"
	"link-shortener/pkg/logger"
	"link-shortener/pkg/middleware"
	"link-shortener/pkg/request"
	"link-shortener/pkg/response"
	"net/http"
)

var log = logger.GetLogger()

type Handler struct {
	LinkHandlerDeps
}

type userService interface {
	FindByEmail(email string) (*model.User, error)
}

type LinkHandlerDeps struct {
	Router      *http.ServeMux
	Service     *Service
	UserService userService
	StatService *stat.Service
	Config      *config.Config
}

func RegisterLinkHandler(linkHandlerDeps LinkHandlerDeps) {
	handler := &Handler{
		linkHandlerDeps,
	}
	handler.Router.Handle("POST /link", middleware.IsAuthed(http.HandlerFunc(handler.create), *handler.Config))
	handler.Router.Handle("GET /{hash}", http.HandlerFunc(handler.goTo))
	handler.Router.Handle("DELETE /link", middleware.IsAuthed(http.HandlerFunc(handler.delete), *handler.Config))
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	userData, err := jwt.GetClaimsFromContext[auth.JwtAuthUserData](r.Context())
	if err != nil {
		apperror.HandleError(apperror.Internal("Failed to proccess JWT token"), w)
		return
	}

	user, err := h.UserService.FindByEmail(userData.Email)
	if err != nil {
		apperror.HandleError(err, w)
		return
	}

	body, err := request.GetBody[CreateRequestPayload](w, r)
	if err != nil {
		apperror.HandleError(apperror.BadRequest(err.Error()), w)
		return
	}

	createdLink, err := h.Service.Create(body.Url, user.ID)
	if err != nil {
		apperror.HandleError(err, w)
		return
	}
	response.Json(w, http.StatusOK, &createdLink)
}

func (h *Handler) goTo(w http.ResponseWriter, r *http.Request) {
	hash := r.PathValue("hash")
	link, err := h.Service.Find(&FindParams{
		hash: hash,
	})
	if err != nil {
		apperror.HandleError(err, w)
		return
	}

	err = h.StatService.AddClick(link.ID)
	if err != nil {
		log.Println("AddClick err ", err.Error())
	}
	http.Redirect(w, r, link.Url, http.StatusSeeOther)
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	userData, err := jwt.GetClaimsFromContext[auth.JwtAuthUserData](r.Context())
	if err != nil {
		apperror.HandleError(apperror.Internal("Failed to proccess JWT token"), w)
		return
	}

	user, err := h.UserService.FindByEmail(userData.Email)
	if err != nil {
		apperror.HandleError(err, w)
		return
	}

	body, err := request.GetBody[DeleteRequestPayload](w, r)
	if err != nil {
		apperror.HandleError(apperror.BadRequest(err.Error()), w)
		return
	}

	err = h.Service.Delete(&FindParams{
		hash:   body.Hash,
		url:    body.Url,
		id:     body.Id,
		userId: user.ID,
	})
	if err != nil {
		apperror.HandleError(err, w)
		return
	}

	response.Json(w, 200, &map[string]string{"message": "Link deleted"})
}
