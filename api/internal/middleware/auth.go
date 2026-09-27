package middleware

import (
	"context"
	"log/slog"
	"net/http"
	"reflect"
	"strings"

	"github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth"
)

type Auth struct {
	log         *slog.Logger
	js          *auth.JwtService
	authService *auth.Service
}

func NewAuth(log *slog.Logger, js *auth.JwtService, authService *auth.Service) *Auth {
	a := Auth{js: js, authService: authService}
	a.log = log.With("module", reflect.TypeOf(a))
	return &Auth{log: log.With(slog.String("module", "middleware.Auth")), js: js, authService: authService}
}

func (a *Auth) TokenToContext(next nethttp.StrictHTTPHandlerFunc, operationID string) nethttp.StrictHTTPHandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, request any) (response any, err error) {
		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer")
		token = strings.TrimSpace(token)
		claims, err := a.js.ValidateToken(token)
		if err != nil {
			// keep ctx as is if no valid token is found
			return next(ctx, w, r, request)
		}
		user, err := a.authService.GetUserByID(claims.Subject)
		newCtx := context.WithValue(ctx, "user", user)
		newCtx = context.WithValue(newCtx, "userId", claims)
		return next(newCtx, w, r, request)
	}
}
