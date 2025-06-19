package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashingService_CheckPassword_IsEqual(t *testing.T) {
	hs := NewHashingService()

	password1 := "that is some password you got there"
	password2 := "that is some password you got there"

	salt := hs.generateSalt()
	hashedPassword := hs.hashPassword(password1, salt)

	assert.True(t, hs.CheckPassword(password2, hashedPassword))
}

func TestAuthService_CheckPassword_IsNotEqual(t *testing.T) {
	hs := NewHashingService()
	password1 := "that is some password you got there"
	password2 := "that is some password you gt there"

	salt := hs.generateSalt()
	hashedPassword := hs.hashPassword(password1, salt)

	assert.False(t, hs.CheckPassword(password2, hashedPassword))
}
