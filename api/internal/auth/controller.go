package auth

import (
	"encoding/json"
	"fmt"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
	"log/slog"
	"net/http"
)

type Controller struct {
	AuthService *Service
}

func NewController(as *Service) *Controller {
	return &Controller{AuthService: as}
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	loginRequest, err := utils.ParseJson[LoginRequest](r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := c.AuthService.LoginUser(loginRequest)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		w.WriteHeader(http.StatusOK)
		enc := json.NewEncoder(w)
		enc.Encode(LoginResponse{Token: token})
	}
}

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	regRequest, err := utils.ParseJson[RegisterRequest](r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		slog.Error("deserializing failed", slog.Any("err", err))
	}

	slog.Info("registering user", slog.Any("request", regRequest))

	err = c.AuthService.RegisterUser(regRequest)

	if err != nil {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintln(w, err.Error())
	}
}
