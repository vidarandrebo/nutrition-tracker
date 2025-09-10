package auth

import (
	"errors"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth/user"
)

type Service struct {
	userStore      user.IStore
	hashingService IHashingService
	jwtService     *JwtService
}

func NewAuthService(store user.IStore, hs IHashingService, jwtService *JwtService) *Service {
	return &Service{
		userStore:      store,
		hashingService: hs,
		jwtService:     jwtService,
	}
}

func (s *Service) RegisterUser(rr Register) error {
	existingUser, _ := s.userStore.GetUserByEmail(rr.Email)
	if existingUser != nil {
		return errors.New("user already exists")
	}

	hash := s.hashingService.HashPassword(rr.Password)

	u := &user.User{
		ID:           0,
		Name:         "",
		Email:        rr.Email,
		PasswordHash: hash,
	}

	err := s.userStore.AddUser(u)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) LoginUser(lr Login) (LoginResult, error) {
	u, err := s.userStore.GetUserByEmail(lr.Email)
	if err != nil {
		return LoginResult{}, errors.New("incorrect credentials")
	}
	if !s.hashingService.CheckPassword(lr.Password, u.PasswordHash) {
		return LoginResult{}, errors.New("incorrect credentials")
	}
	token, err := s.jwtService.CreateToken(u.ID)
	return LoginResult{token, u.ID}, err
}

type LoginResult struct {
	Token  string
	UserID int64
}
