package meal

import (
	"database/sql"
	"log/slog"
)

type Store struct {
	DB     *sql.DB
	Logger *slog.Logger
}

func NewStore(db *sql.DB, logger *slog.Logger) *Store {
	return &Store{
		DB:     db,
		Logger: logger,
	}
}
