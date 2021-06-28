package users

import (
	"context"
	ts "users.org/pkg/time"
)

// UserRepositoryMem is a in memory repository implementation for User model
// Implements UserRepository
type UserRepositoryMem struct {
	users map[string]*User // [user_id]user
}

// NewUserRepositoryMem creates new UserRepositoryMem instance
func NewUserRepositoryMem() *UserRepositoryMem {
	return &UserRepositoryMem{
		users: make(map[string]*User),
	}
}

// Insert appends new user into repository.
func (r *UserRepositoryMem) Insert(ctx context.Context, user *User) error {
	user.ID = generateUserID()
	user.RecordingDate = ts.Now()

	r.users[user.ID] = user

	return nil
}

// ReadByID returns items from repository by given ID
// Returns nil if there is no user with provided ID
func (r *UserRepositoryMem) ReadByID(ctx context.Context, id string) (*User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, nil
	}

	return user, nil
}

// ReadAll returns all items from repository
func (r *UserRepositoryMem) ReadAll(ctx context.Context) ([]*User, error) {
	items := make([]*User, 0, len(r.users))

	for _, user := range r.users {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()

		default:
			items = append(items, user)
		}
	}

	return items, nil
}
