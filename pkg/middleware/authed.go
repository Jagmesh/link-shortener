package middleware

import (
	"context"
	"link-shortener/config"
	apperror "link-shortener/pkg/app-error"
	"link-shortener/pkg/jwt"
	"net/http"
	"strings"
)

func IsAuthed(next http.Handler, config config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			apperror.HandleError(apperror.Unauthorized("Pass a valid token in 'Authorization' header"), w)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := jwt.NewJWT(config.Auth.Secret).Parse(token)
		if err != nil || claims["email"] == "" {
			apperror.HandleError(apperror.Forbidden("Invalid token"), w)
			return
		}

		ctx := context.WithValue(r.Context(), jwt.CLAIMS_CTX_KEY, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
