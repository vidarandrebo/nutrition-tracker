package auth

import (
	"net/mail"
)

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (rr RegisterRequest) Validate() bool {
	if _, err := mail.ParseAddress(rr.Email); err != nil {
		return false
	}
	if len(rr.Password) < 8 {
		return false
	}
	return true
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
