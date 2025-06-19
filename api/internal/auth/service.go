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

func (s *Service) RegisterUser(rr RegisterRequest) error {
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

	s.userStore.AddUser(u)
	return nil
}

func (s *Service) LoginUser(lr LoginRequest) (string, error) {
	user, err := s.userStore.GetUserByEmail(lr.Email)
	if err != nil {
		return "", errors.New("incorrect credentials")
	}
	if s.hashingService.CheckPassword(lr.Password, user.PasswordHash) {
		token, err := s.jwtService.CreateToken(user.ID)
		return token, err
	}
	return "", errors.New("incorrect credentials")
}
