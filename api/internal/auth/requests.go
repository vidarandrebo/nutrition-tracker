package auth

import (
	"net/mail"
)

type Register struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (rr Register) Validate() bool {
	if _, err := mail.ParseAddress(rr.Email); err != nil {
		return false
	}
	if len(rr.Password) < 8 {
		return false
	}
	return true
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
