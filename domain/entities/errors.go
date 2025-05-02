package entities

import "errors"

var (
	// User domain errors
	ErrInvalidName        = errors.New("user: name cannot be empty")
	ErrInvalidPassword    = errors.New("user: password hash cannot be empty")
	ErrAlreadyBlocked     = errors.New("user: already blocked")
	ErrNotBlocked         = errors.New("user: not blocked")
	ErrCannotDeleteActive = errors.New("user: cannot delete active user")

	// Task domain errors
	ErrEmptyTaskName  = errors.New("task: name cannot be empty")
	ErrDeadlineInPast = errors.New("task: deadline must be in the future")
	ErrAlreadyDone    = errors.New("task: already marked as done")
)
