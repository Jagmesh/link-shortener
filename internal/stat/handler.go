package stat

import (
	"link-shortener/config"
	"link-shortener/entity/model"
	"link-shortener/internal/auth"
	"link-shortener/internal/link"
	apperror "link-shortener/pkg/app-error"
	"link-shortener/pkg/jwt"
	"link-shortener/pkg/middleware"
	"link-shortener/pkg/request"
	"link-shortener/pkg/response"
	"net/http"
)

type userService interface {
	Create(email, password, name string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}

type Handler struct {
	*HandlerDeps
}

type HandlerDeps struct {
	Router      *http.ServeMux
	Config      *config.Config
	Jwt         *jwt.JWT
	Service     *Service
	UserService userService
	LinkService *link.Service
}

func RegisterStatHandler(deps *HandlerDeps) {
	handler := &Handler{
		deps,
	}
	handler.Router.Handle("GET /stat", middleware.IsAuthed(http.HandlerFunc(handler.getStats), *handler.Config))
}

func (h *Handler) getStats(w http.ResponseWriter, r *http.Request) {
	queryParams, err := request.GetQueryParams[GetRequestPayload](r)
	if err != nil {
		apperror.HandleError(apperror.BadRequest(err.Error()), w)
		return
	}

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

	foundLinks, err := h.LinkService.FindAll(&link.FindParams{UserId: user.ID})
	var linksId []uint
	for _, foundLink := range foundLinks {
		linksId = append(linksId, foundLink.ID)
	}

	clicks, err := h.Service.GetClicksNumberByDate(linksId, queryParams.From, queryParams.To)
	if err != nil {
		apperror.HandleError(err, w)
		return
	}

	response.Json(w, 200, &map[string]uint{"clicks": clicks})
}
