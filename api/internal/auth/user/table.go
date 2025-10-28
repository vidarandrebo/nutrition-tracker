package user

import "time"

type TableUser struct {
	ID           int64
	Name         string
	Email        string
	PasswordHash []byte
	DateCreated  time.Time
	DateModified time.Time
}
