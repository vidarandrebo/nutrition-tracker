package auth

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth/user"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
)

type failingUserRepository struct {
	user.IRepository
	getErr error
	t      *testing.T
}

func (r failingUserRepository) GetByEmail(string) (user.TableUser, error) {
	return user.TableUser{}, r.getErr
}

func (r failingUserRepository) Add(user.TableUser) (user.TableUser, error) {
	r.t.Fatal("unexpected add after lookup failure")
	return user.TableUser{}, nil
}

func TestAuthServiceLookupFailure(t *testing.T) {
	service := NewAuthService(failingUserRepository{getErr: errors.New("database unavailable"), t: t}, nil, nil)
	assert.ErrorIs(t, service.RegisterUser(Register{}), utils.ErrUnknown)
	_, err := service.LoginUser(Login{})
	assert.ErrorIs(t, err, utils.ErrUnknown)
}

func TestAuthServiceMissingUser(t *testing.T) {
	service := NewAuthService(failingUserRepository{getErr: sql.ErrNoRows, t: t}, nil, nil)
	_, err := service.LoginUser(Login{})
	assert.ErrorIs(t, err, utils.ErrUnauthorized)
}

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
