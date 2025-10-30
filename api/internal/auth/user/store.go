package user

import (
	"database/sql"
	"log/slog"
	"reflect"

	"github.com/vidarandrebo/nutrition-tracker/api/internal/utils"
)

type IRepository interface {
	Add(item TableUser) (TableUser, error)
	GetByEmail(email string) (TableUser, error)
}

type Repository struct {
	db  *sql.DB
	log *slog.Logger
}

func NewRepository(db *sql.DB, log *slog.Logger) *Repository {
	r := Repository{
		db: db,
	}
	r.log = log.With(slog.Any("module", reflect.TypeOf(r)))
	return &r
}

func (s *Repository) Add(item TableUser) (TableUser, error) {
	scanErr := s.db.QueryRow(`
		INSERT INTO users(name, email, password_hash) 
		VALUES ($1, $2, $3)`,
		item.Name, item.Email, item.PasswordHash,
	).Scan(&item.ID)

	if scanErr != nil {
		return TableUser{}, utils.ErrUnknown
	}

	return item, nil
}

func (s *Repository) GetByEmail(email string) (TableUser, error) {
	user := TableUser{}
	scanErr := s.db.QueryRow(`
		SELECT id, name, email, password_hash 
		FROM users AS u 
		WHERE u.email=$1`,
		email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
	)

	if scanErr != nil {
		s.log.Error("no user matching the credentials", slog.String("email", email))
		return TableUser{}, utils.ErrEntityNotFound
	}

	return user, nil
}
