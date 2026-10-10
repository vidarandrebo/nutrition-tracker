package utils

import "errors"

var (

	// ErrNotFound represents an error indicating that the requested resource was not found in the system.
	ErrNotFound = errors.New("not found")

	// ErrForbidden represents an error indicating that access to the requested resource is forbidden due to insufficient permissions.
	ErrForbidden = errors.New("forbidden")

	// ErrUnauthorized represents an error indicating that valid authentication credentials are required.
	ErrUnauthorized = errors.New("unauthorized")

	// ErrUnknown represents an error indicating that an unknown error occurred.
	ErrUnknown = errors.New("unknown error")
)
