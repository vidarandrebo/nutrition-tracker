package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"strings"

	"github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
)

type Auth struct {
	log *slog.Logger
	js  *auth.JwtService
}

func NewAuth(log *slog.Logger, js *auth.JwtService) *Auth {
	return &Auth{log: log.With(slog.String("module", "middleware.Auth")), js: js}
}

func (a *Auth) TokenToContext(next nethttp.StrictHTTPHandlerFunc, operationID string) nethttp.StrictHTTPHandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (response interface{}, err error) {
		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer")
		token = strings.TrimSpace(token)
		claims, err := a.js.ValidateToken(token)

		if err != nil {
			// keep ctx as is if no valid token is found
			a.log.Warn("authentication failure", slog.Any("error", err))
			return next(ctx, w, r, request)
		} else {
			newCtx := context.WithValue(ctx, "user", claims)
			return next(newCtx, w, r, request)
		}
	}
}
