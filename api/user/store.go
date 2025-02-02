package user

import (
	"database/sql"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) AddUser(user *User) {
	_, err := s.db.Exec("INSERT INTO users(name, passwordhash) VALUES ($1, $2)", user.Name, user.PasswordHash)
	if err != nil {
		panic(err)
	}
}

func (s *Store) ListUsers() []*User {
	users := make([]*User, 0)
	rows, err := s.db.Query("SELECT id, name, passwordhash FROM users")

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
func (s *Store) GetUser(email string) *User {
	row := s.db.QueryRow("SELECT id, name, passwordhash FROM users")

	user := User{}
	err := row.Scan(&user.ID, &user.Name, &user.PasswordHash)

	if err != nil {
		panic(err)
	}

	return &user
}
