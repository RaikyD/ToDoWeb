package entities

import "github.com/google/uuid"

// User represents an application user with credentials and state.
type User struct {
	ID       uuid.UUID
	Name     string // login or display name
	Password string // hashed password
	Blocked  bool   // account blocked state
}

// NewUser creates a new User with a generated ID and initial state.
// Returns ErrInvalidName or ErrInvalidPassword for invalid input.
func NewUser(name, passwordHash string) (*User, error) {
	if name == "" {
		return nil, ErrInvalidName
	}
	if passwordHash == "" {
		return nil, ErrInvalidPassword
	}
	return &User{
		ID:       uuid.New(),
		Name:     name,
		Password: passwordHash,
		Blocked:  false,
	}, nil
}

// Block marks the user as blocked. Returns ErrAlreadyBlocked if already blocked.
func (u *User) Block() error {
	if u.Blocked {
		return ErrAlreadyBlocked
	}
	u.Blocked = true
	return nil
}

// Unblock marks the user as unblocked. Returns ErrNotBlocked if not currently blocked.
func (u *User) Unblock() error {
	if !u.Blocked {
		return ErrNotBlocked
	}
	u.Blocked = false
	return nil
}

// CanDelete validates whether the user can be deleted according to business rules.
// For example, only blocked users can be deleted.
func (u *User) CanDelete() error {
	if !u.Blocked {
		return ErrCannotDeleteActive
	}
	return nil
}
