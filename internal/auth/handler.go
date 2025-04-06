package auth

import (
	"link-shortener/config"
	apperror "link-shortener/pkg/app-error"
	"link-shortener/pkg/jwt"
	"link-shortener/pkg/request"
	"link-shortener/pkg/response"
	"net/http"
)

type authHandler struct {
	deps AuthHandlerDeps
}

type AuthHandlerDeps struct {
	Router      *http.ServeMux
	Config      *config.Config
	Jwt         *jwt.JWT
	AuthService *Service
}

func ReqisterAuthHandler(AuthHandlerDeps AuthHandlerDeps) {
	handler := &authHandler{
		deps: AuthHandlerDeps,
	}
	handler.deps.Router.HandleFunc("POST /auth/login", handler.login)
	handler.deps.Router.HandleFunc("POST /auth/register", handler.register)
}

func (h *authHandler) login(w http.ResponseWriter, r *http.Request) {
	body, err := request.GetBody[LoginRequestPayload](w, r)
	if err != nil {
		apperror.HandleError(apperror.BadRequest(err.Error()), w)
		return
	}

	user, err := h.deps.AuthService.Login(body.Email, body.Password)
	if err != nil {
		apperror.HandleError(err, w)
		return
	}

	token, err := h.deps.Jwt.Create(map[string]any{
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

func (h *authHandler) register(w http.ResponseWriter, r *http.Request) {
	body, err := request.GetBody[RegisterRequestPayload](w, r)
	if err != nil {
		apperror.HandleError(apperror.BadRequest(err.Error()), w)
		return
	}

	user, err := h.deps.AuthService.Register(body.Email, body.Password, body.Name)
	if err != nil {
		apperror.HandleError(apperror.BadRequest(err.Error()), w)
		return
	}

	token, err := h.deps.Jwt.Create(map[string]any{
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
