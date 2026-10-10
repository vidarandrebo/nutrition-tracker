package auth

import (
	"database/sql"
	"errors"
	"log/slog"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/auth/user"
	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
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
	_, byEmailErr := s.userRepository.GetByEmail(rr.Email)
	if byEmailErr == nil {
		return errors.New("user already exists")
	}
	if !errors.Is(byEmailErr, sql.ErrNoRows) {
		slog.Error("failed to check existing user during registration", slog.Any("err", byEmailErr))
		return utils.ErrUnknown
	}

	hash := s.hashingService.HashPassword(rr.Password)

	u := user.TableUser{
		ID:           0,
		Name:         "",
		Email:        rr.Email,
		PasswordHash: hash,
	}

	_, err := s.userRepository.Add(u)
	if err != nil {
		slog.Error("failed to register user", slog.Any("err", err))
		return utils.ErrUnknown
	}
	return nil
}

func (s *Service) LoginUser(lr Login) (LoginResult, error) {
	u, err := s.userRepository.GetByEmail(lr.Email)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			slog.Error("failed to look up user during login", slog.Any("err", err))
			return LoginResult{}, utils.ErrUnknown
		}
		slog.Debug("login user not found", slog.Any("err", err))
		return LoginResult{}, utils.ErrUnauthorized
	}
	if !s.hashingService.CheckPassword(lr.Password, u.PasswordHash) {
		return LoginResult{}, utils.ErrUnauthorized
	}
	token, err := s.jwtService.CreateToken(u.ID)
	if err != nil {
		slog.Error("failed to create login token", slog.Int64("userID", u.ID), slog.Any("err", err))
		return LoginResult{}, utils.ErrUnknown
	}
	return LoginResult{token, u.ID}, nil
}

func (s *Service) GetUserByID(id int64) (*user.User, error) {
	u, err := s.userRepository.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Debug("user not found", slog.Int64("userID", id), slog.Any("err", err))
			return nil, utils.ErrNotFound
		}
		slog.Error("failed to get user", slog.Int64("userID", id), slog.Any("err", err))
		return nil, utils.ErrUnknown
	}
	user := user.NewUser()
	user.FromTable(u)
	return user, nil
}

type LoginResult struct {
	Token  string
	UserID int64
}
