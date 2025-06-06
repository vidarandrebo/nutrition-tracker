package user

import (
	"database/sql"
	"errors"
	"log/slog"
)

type IStore interface {
	AddUser(user *User)
	ListUsers() []*User
	GetUserByEmail(string) (*User, error)
}
type Store struct {
	db  *sql.DB
	log *slog.Logger
}

func NewStore(db *sql.DB, log *slog.Logger) *Store {
	return &Store{
		db:  db,
		log: log.With(slog.String("module", "user.Store")),
	}
}

func (s *Store) AddUser(user *User) {
	_, err := s.db.Exec(`
		INSERT INTO users(name, email, password_hash) 
		VALUES ($1, $2, $3)`,
		user.Name, user.Email, user.PasswordHash)
	if err != nil {
		panic(err)
	}
}

func (s *Store) ListUsers() []*User {
	users := make([]*User, 0)
	rows, err := s.db.Query("SELECT id, name, password_hash FROM users")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		user := User{}
		rows.Scan(&user.ID, &user.Name, &user.PasswordHash)
		users = append(users, &user)
	}

	return users
}
func (s *Store) GetUserByEmail(email string) (*User, error) {
	row := s.db.QueryRow(`
		SELECT id, name, email, password_hash 
		FROM users as u 
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
