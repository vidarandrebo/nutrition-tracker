package utils

import "errors"

var (
	ErrEntityNotFound     = errors.New("entity not found")
	ErrEntityNotOwned     = errors.New("entity not owned")
	ErrEntityInaccessible = errors.New("entity not accessible")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrUnknown            = errors.New("an unknown error occurred")
)
