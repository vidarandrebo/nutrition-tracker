package user

import (
	"database/sql"
	"errors"
	"log/slog"
	"reflect"
)

type IStore interface {
	AddUser(user *User) error
	ListUsers() ([]*User, error)
	GetUserByEmail(string) (*User, error)
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

func (s *Repository) AddUser(user *User) error {
	_, err := s.db.Exec(`
		INSERT INTO users(name, email, password_hash) 
		VALUES ($1, $2, $3)`,
		user.Name, user.Email, user.PasswordHash)
	if err != nil {
		panic(err)
		return err
	}
	return nil
}

func (s *Repository) ListUsers() ([]*User, error) {
	users := make([]*User, 0)
	rows, err := s.db.Query("SELECT id, name, password_hash FROM users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := User{}
		rows.Scan(&user.ID, &user.Name, &user.PasswordHash)
		users = append(users, &user)
	}

	return users, nil
}

func (s *Repository) GetUserByEmail(email string) (*User, error) {
	row := s.db.QueryRow(`
		SELECT id, name, email, password_hash 
		FROM users AS u 
		WHERE u.email=$1`,
		email)

	user := User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash)
	if err != nil {
		s.log.Info("no user matching the credentials", slog.String("email", email))
		return nil, errors.New("no user matching the email")
	}

	return &user, nil
}
