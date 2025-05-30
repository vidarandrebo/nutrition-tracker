package recipe

import (
	"database/sql"
	"log/slog"
	"reflect"
)

type Store struct {
	logger *slog.Logger
	DB     *sql.DB
}

func NewStore(db *sql.DB, logger *slog.Logger) *Store {
	s := Store{DB: db}
	s.logger = logger.With(slog.Any("module", reflect.TypeOf(s)))
	return &s
}

func (s *Store) Get(ownerID int64) ([]Recipe, error) {
	recipes := make([]Recipe, 0)
	return recipes, nil
}

func (s *Store) Add(recipe Recipe, ownerID int64) (Recipe, error) {
	return Recipe{}, nil
}
