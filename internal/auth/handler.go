package auth

import (
	"link-shortener/config"
	apperror "link-shortener/pkg/app-error"
	"link-shortener/pkg/jwt"
	"link-shortener/pkg/request"
	"link-shortener/pkg/response"
	"net/http"
)

type Handler struct {
	*HandlerDeps
}

type HandlerDeps struct {
	Router      *http.ServeMux
	Config      *config.Config
	Jwt         *jwt.JWT
	AuthService *Service
}

func RegisterAuthHandler(AuthHandlerDeps *HandlerDeps) {
	handler := &Handler{
		AuthHandlerDeps,
	}
	handler.Router.HandleFunc("POST /auth/login", handler.login)
	handler.Router.HandleFunc("POST /auth/register", handler.register)
}

func (h *Handler) login(w http.ResponseWriter, r *http.Request) {
	body, err := request.GetBody[LoginRequestPayload](r)
	if err != nil {
		apperror.HandleError(apperror.BadRequest(err.Error()), w)
		return
	}

	user, err := h.AuthService.Login(body.Email, body.Password)
	if err != nil {
		apperror.HandleError(err, w)
		return
	}

	token, err := h.Jwt.Create(map[string]any{
		"email": user.Email,
	})
	if err != nil {
		apperror.HandleError(apperror.Unauthorized(err.Error()), w)
		return
	}

	response.Json(w, 200, &LoginResponsePayload{
		Token: token,
	})
}

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {
	body, err := request.GetBody[RegisterRequestPayload](r)
	if err != nil {
		apperror.HandleError(apperror.BadRequest(err.Error()), w)
		return
	}

	user, err := h.AuthService.Register(body.Email, body.Password, body.Name)
	if err != nil {
		apperror.HandleError(apperror.BadRequest(err.Error()), w)
		return
	}

	token, err := h.Jwt.Create(map[string]any{
		"email": user.Email,
	})
	if err != nil {
		apperror.HandleError(apperror.BadRequest(err.Error()), w)
		return
	}

	response.Json(w, 200, &LoginResponsePayload{
		Token: token,
	})
}
