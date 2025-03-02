package auth

import (
	"fmt"
	"link-shortener/config"
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

type Test struct {
	SomeField string
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
		response.Json(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	fmt.Println(body)

	response.Json(w, 200, &LoginResponsePayload{
		Token: "token",
	})
}

func (h *authHandler) register(w http.ResponseWriter, r *http.Request) {
	body, err := request.GetBody[RegisterRequestPayload](w, r)
	if err != nil {
		response.Json(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	user, err := h.deps.AuthService.Register(body.Email, body.Password, body.Name)
	if err != nil {
		response.Json(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	token, err := h.deps.Jwt.Create(map[string]any{
		"email": user.Email,
	})
	if err != nil {
		response.Json(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	claims, err := h.deps.Jwt.Parse(token)
	fmt.Println("err", err)
	fmt.Println("claims", claims)

	response.Json(w, 200, &LoginResponsePayload{
		Token: token,
	})
}
