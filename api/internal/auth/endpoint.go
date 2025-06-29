package auth

import (
	"context"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/api"
)

type Endpoint struct {
}

func (e Endpoint) PostApiLogin(ctx context.Context, request api.PostApiLoginRequestObject) (api.PostApiLoginResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (e Endpoint) PostApiRegister(ctx context.Context, request api.PostApiRegisterRequestObject) (api.PostApiRegisterResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
