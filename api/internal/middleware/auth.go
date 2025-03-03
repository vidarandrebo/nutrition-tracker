package middleware

import (
	"context"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
	"log/slog"
	"net/http"
	"strings"
)

type Auth struct {
	log *slog.Logger
	js  *auth.JwtService
}

func NewAuth(log *slog.Logger, js *auth.JwtService) *Auth {
	return &Auth{log: log.With(slog.String("module", "middleware.Auth")), js: js}
}

func (rt *Auth) TokenToContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer")
		token = strings.TrimSpace(token)
		claims, err := rt.js.ValidateToken(token)

		if err != nil {
			// keep ctx as is if no valid token is found
			rt.log.Warn("authentication failure", slog.Any("error", err))
			next.ServeHTTP(w, r)
		} else {
			newCtx := context.WithValue(r.Context(), "user", claims)
			next.ServeHTTP(w, r.WithContext(newCtx))
		}
	})
}
