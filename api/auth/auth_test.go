package auth

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthService_CheckPassword_IsEqual(t *testing.T) {
	as := NewAuthService()

	password1 := "that is some password you got there"
	password2 := "that is some password you got there"

	salt := as.GenerateSalt()
	hashedPassword := as.HashPassword(password1, salt)

	assert.True(t, as.CheckPassword(password2, hashedPassword))
}
func TestAuthService_CheckPassword_IsNotEqual(t *testing.T) {
	as := NewAuthService()

	password1 := "that is some password you got there"
	password2 := "that is some password you gt there"

	salt := as.GenerateSalt()
	hashedPassword := as.HashPassword(password1, salt)

	assert.False(t, as.CheckPassword(password2, hashedPassword))
}
