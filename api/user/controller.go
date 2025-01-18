package user

import (
	"encoding/json"
	"fmt"
	"github.com/vidarandrebo/nutrition-tracker/api/auth"
	"io"
	"log/slog"
	"net/http"
)

type Controller struct {
	Store       *Store
	AuthService *auth.AuthService
}

func NewController(store *Store) *Controller {
	return &Controller{Store: store}
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		panic(err)
	}
	user := c.Store.GetUser(credentials.Email)
	if c.AuthService.CheckPassword(credentials.Password, user.PasswordHash) {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	credentials, err := ParseJson[Credentials](r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		slog.Error("deserializing failed", slog.Any("err", err))
	}
	fmt.Fprintln(w, credentials)
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ParseJson[T any](reader io.Reader) (*T, error) {
	var data T
	err := json.NewDecoder(reader).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
