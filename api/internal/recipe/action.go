package recipe

import (
	"log/slog"
	"net/http"
	"reflect"
)

type ActionBase struct {
	logger *slog.Logger
	store  *Store
}

func NewAction(store *Store, logger *slog.Logger) *ActionBase {
	a := ActionBase{
		store: store,
	}
	a.logger = logger.With(slog.Any("module", reflect.TypeOf(a)))
	return &a
}

type Actions struct {
	Post    http.Handler
	GetById http.Handler
	Get     http.Handler
	Put     http.Handler
	Delete  http.Handler
}

func NewActions(store *Store, logger *slog.Logger) *Actions {
	actionBase := NewAction(store, logger)
	return &Actions{
		Post:    NewPost(actionBase),
		GetById: nil,
		Get:     nil,
		Put:     nil,
		Delete:  nil,
	}
}
