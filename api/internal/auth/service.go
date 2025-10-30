package auth

import (
	"errors"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth/user"
)

type Service struct {
	userRepository user.IRepository
	hashingService IHashingService
	jwtService     *JwtService
}

func NewAuthService(userRepository user.IRepository, hs IHashingService, jwtService *JwtService) *Service {
	return &Service{
		userRepository: userRepository,
		hashingService: hs,
		jwtService:     jwtService,
	}
}

func (s *Service) RegisterUser(rr Register) error {
	_, err := s.userRepository.GetByEmail(rr.Email)
	if err == nil {
		return errors.New("user already exists")
	}

	hash := s.hashingService.HashPassword(rr.Password)

	u := user.TableUser{
		ID:           0,
		Name:         "",
		Email:        rr.Email,
		PasswordHash: hash,
	}

	_, err = s.userRepository.Add(u)

	return err
}

func (s *Service) LoginUser(lr Login) (LoginResult, error) {
	u, err := s.userRepository.GetByEmail(lr.Email)
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
