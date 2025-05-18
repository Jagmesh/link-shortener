package link

import (
	"link-shortener/config"
	"link-shortener/entity/event"
	"link-shortener/entity/model"
	"link-shortener/internal/auth"
	apperror "link-shortener/pkg/app-error"
	"link-shortener/pkg/bus"
	"link-shortener/pkg/jwt"
	"link-shortener/pkg/middleware"
	"link-shortener/pkg/request"
	"link-shortener/pkg/response"
	"net/http"
)

type Handler struct {
	*HandlerDeps
}

type userService interface {
	FindByEmail(email string) (*model.User, error)
}

type HandlerDeps struct {
	Router      *http.ServeMux
	Service     *Service
	UserService userService
	Config      *config.Config
	EventBus    *bus.EventBus
}

func RegisterLinkHandler(deps *HandlerDeps) {
	handler := &Handler{
		deps,
	}
	handler.Router.Handle("POST /link", middleware.IsAuthed(http.HandlerFunc(handler.create), *handler.Config))
	handler.Router.Handle("GET /{hash}", http.HandlerFunc(handler.goTo))
	handler.Router.Handle("DELETE /link", middleware.IsAuthed(http.HandlerFunc(handler.delete), *handler.Config))
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	userData, err := jwt.GetClaimsFromContext[auth.JwtAuthUserData](r.Context())
	if err != nil {
		apperror.HandleError(apperror.Internal("Failed to process JWT token"), w)
		return
	}

	user, err := h.UserService.FindByEmail(userData.Email)
	if err != nil {
		apperror.HandleError(err, w)
		return
	}

	body, err := request.GetBody[CreateRequestPayload](r)
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
	link, err := h.Service.FindOne(&FindParams{
		Hash: hash,
	})
	if err != nil {
		apperror.HandleError(err, w)
		return
	}

	go h.EventBus.Publish(event.NewStatClickEvent(link.ID))

	http.Redirect(w, r, link.Url, http.StatusSeeOther)
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request) {
	userData, err := jwt.GetClaimsFromContext[auth.JwtAuthUserData](r.Context())
	if err != nil {
		apperror.HandleError(apperror.Internal("Failed to process JWT token"), w)
		return
	}

	user, err := h.UserService.FindByEmail(userData.Email)
	if err != nil {
		apperror.HandleError(err, w)
		return
	}

	body, err := request.GetBody[DeleteRequestPayload](r)
	if err != nil {
		apperror.HandleError(apperror.BadRequest(err.Error()), w)
		return
	}

	err = h.Service.Delete(&FindParams{
		Hash:   body.Hash,
		Url:    body.Url,
		Id:     body.Id,
		UserId: user.ID,
	})
	if err != nil {
		apperror.HandleError(err, w)
		return
	}

	response.Json(w, 200, &map[string]string{"message": "Link deleted"})
}
