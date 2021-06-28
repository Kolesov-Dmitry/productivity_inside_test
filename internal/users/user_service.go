package users

import (
	"context"
)

// UserService provides high level API to work with users
type UserService struct {
	repo UserRepository
}

// NewUserService creates new UserService instance
func NewUserService(r UserRepository) *UserService {
	return &UserService{
		repo: r,
	}
}

// AddUser saves new user
func (s *UserService) AddUser(ctx context.Context, user *User) error {
	return s.repo.Insert(ctx, user)
}

// GetUser returns user by given ID. Returns all users if ID is null
func (s *UserService) GetUser(ctx context.Context, id string) ([]*User, error) {
	if len(id) == 0 {
		return s.repo.ReadAll(ctx)
	}

	user, err := s.repo.ReadByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return []*User{user}, nil
}
