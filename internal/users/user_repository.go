package users

import (
	"context"
)

// UserRepository is a repository abstraction for User model
type UserRepository interface {
	// Insert appends new user into repository.
	Insert(ctx context.Context, user *User) error

	// ReadByID returns items from repository by given ID
	// Returns nil if there is no user with provided ID
	ReadByID(ctx context.Context, id string) (*User, error)

	// ReadAll returns all items from repository
	ReadAll(ctx context.Context) ([]*User, error)
}
