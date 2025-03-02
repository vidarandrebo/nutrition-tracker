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
	Logger      *slog.Logger
}

func NewController(as *Service, log *slog.Logger) *Controller {
	return &Controller{AuthService: as, Logger: log.With(slog.String("module", "auth.Controller"))}
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
		c.Logger.Error("deserializing failed", slog.Any("err", err))
		return
	}
	if !regRequest.Validate() {
		w.WriteHeader(http.StatusBadRequest)
		c.Logger.Error("credential validation failed")
		return
	}

	err = c.AuthService.RegisterUser(regRequest)

	if err != nil {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintln(w, err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
	}

}
