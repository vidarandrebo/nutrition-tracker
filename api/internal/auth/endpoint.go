package auth

import (
	"context"
	"errors"
	"log/slog"
	"reflect"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/api"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
)

type Endpoint struct {
	AuthService *Service
	Logger      *slog.Logger
}

func NewEndpoint(authService *Service, logger *slog.Logger) *Endpoint {
	e := Endpoint{AuthService: authService}
	e.Logger = logger.With("module", reflect.TypeOf(e))
	return &e
}

func (e Endpoint) PostApiLogin(ctx context.Context, request api.PostApiLoginRequestObject) (api.PostApiLoginResponseObject, error) {
	result, err := e.AuthService.LoginUser(Login{
		Email:    request.Body.Email,
		Password: request.Body.Password,
	})
	if err != nil {
		return nil, err
	}
	return api.PostApiLogin200JSONResponse(api.LoginResponse{Token: result.Token, Id: result.UserID}), nil
}

func (e Endpoint) PostApiRegister(ctx context.Context, request api.PostApiRegisterRequestObject) (api.PostApiRegisterResponseObject, error) {
	regRequest := Register{
		Email:    request.Body.Email,
		Password: request.Body.Password,
	}
	if !regRequest.Validate() {
		return nil, errors.New("invalid model")
	}

	err := e.AuthService.RegisterUser(regRequest)
	if err != nil {
		e.Logger.Error("failed to register user", slog.Any("err", err))
		return nil, utils.ErrUnknown
	}
	return api.PostApiRegister201Response{}, nil
}
