package user

import (
	"database/sql"
	"log/slog"
	"reflect"
)

type IRepository interface {
	Add(item TableUser) (TableUser, error)
	GetByEmail(email string) (TableUser, error)
	GetByID(id int64) (TableUser, error)
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

// Add returns sql.ErrNoRows if the insert returns no row, or the database error
// if inserting or scanning the user fails.
func (s *Repository) Add(item TableUser) (TableUser, error) {
	scanErr := s.db.QueryRow(`
		INSERT INTO users(name, email, password_hash) 
		VALUES ($1, $2, $3)
		RETURNING id
		`,
		item.Name, item.Email, item.PasswordHash,
	).Scan(&item.ID)

	if scanErr != nil {
		s.log.Error("failed to add user", slog.String("email", item.Email), slog.Any("error", scanErr))
		return TableUser{}, scanErr
	}

	return item, nil
}

// GetByEmail returns sql.ErrNoRows if the user does not exist, or the database
// error if querying or scanning it fails.
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
		s.log.Error("failed to get user by email", slog.String("email", email), slog.Any("err", scanErr))
		return TableUser{}, scanErr
	}

	return user, nil
}

// GetByID returns sql.ErrNoRows if the user does not exist, or the database
// error if querying or scanning it fails.
func (s *Repository) GetByID(id int64) (TableUser, error) {
	user := TableUser{}
	scanErr := s.db.QueryRow(`
		SELECT id, name, email, password_hash 
		FROM users AS u 
		WHERE u.id=$1`,
		id,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
	)

	if scanErr != nil {
		s.log.Error("failed to get user by id", slog.Int64("id", id), slog.Any("err", scanErr))
		return TableUser{}, scanErr
	}

	return user, nil
}
