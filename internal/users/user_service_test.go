package users_test

import (
	"context"
	"reflect"
	"testing"
	"users.org/internal/users"
	"users.org/mocks"
	ts "users.org/pkg/time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Expected constants
const (
	expectedUserID = "12345"
)

var (
	expectedUsers = []*users.User{
		{
			ID:            "123",
			Name:          "User 1",
			LastName:      "Test",
			Age:           30,
			RecordingDate: ts.Now(),
		},
		{
			ID:            "456",
			Name:          "User 2",
			LastName:      "Test",
			Age:           30,
			RecordingDate: ts.Now(),
		},
	}
)

func makeUserRepositoryMock() *mocks.UserRepository {
	repoMock := &mocks.UserRepository{}

	repoMock.On("Insert", mock.Anything, mock.AnythingOfType("*users.User")).Return(
		func(ctx context.Context, user *users.User) error {
			user.ID = expectedUserID
			return nil
		},
	)

	repoMock.On("ReadByID", mock.Anything, mock.AnythingOfType("string")).Return(
		func(ctx context.Context, id string) *users.User {
			return &users.User{
				ID:            id,
				Name:          "User",
				LastName:      "Test",
				Age:           30,
				RecordingDate: ts.Now(),
			}
		},

		func(ctx context.Context, id string) error {
			return nil
		},
	)

	repoMock.On("ReadAll", mock.Anything).Return(
		func(ctx context.Context) []*users.User {
			return expectedUsers
		},

		func(ctx context.Context) error {
			return nil
		},
	)

	return repoMock
}

func TestUserService_AddUser(t *testing.T) {
	// Setup UserService
	repoMock := makeUserRepositoryMock()
	userService := users.NewUserService(repoMock)

	// Setup Test User
	testUser := users.User{
		Name:          "User 1",
		LastName:      "Test",
		Age:           30,
		RecordingDate: ts.Now(),
	}

	// Call AddUser
	err := userService.AddUser(context.Background(), &testUser)

	// Checks
	assert.Nil(t, err)
	assert.Equal(t, expectedUserID, testUser.ID)

	repoMock.AssertNumberOfCalls(t, "Insert", 1)
	repoMock.AssertNumberOfCalls(t, "ReadByID", 0)
	repoMock.AssertNumberOfCalls(t, "ReadAll", 0)
}

func TestUserService_GetUser_WithID(t *testing.T) {
	// Setup UserService
	repoMock := makeUserRepositoryMock()
	userService := users.NewUserService(repoMock)

	// Call GetUser with ID
	usersList, err := userService.GetUser(context.Background(), expectedUserID)

	// Checks
	assert.Nil(t, err)
	assert.NotNil(t, usersList)
	assert.Equal(t, 1, len(usersList))
	assert.Equal(t, expectedUserID, usersList[0].ID)

	repoMock.AssertNumberOfCalls(t, "ReadByID", 1)
	repoMock.AssertNumberOfCalls(t, "Insert", 0)
	repoMock.AssertNumberOfCalls(t, "ReadAll", 0)
}

func TestUserService_GetUser_WithoutID(t *testing.T) {
	// Setup UserService
	repoMock := makeUserRepositoryMock()
	userService := users.NewUserService(repoMock)

	// Call GetUser without ID
	usersList, err := userService.GetUser(context.Background(), "")

	// Checks
	assert.Nil(t, err)
	assert.NotNil(t, usersList)
	assert.True(t, reflect.DeepEqual(usersList, expectedUsers))

	repoMock.AssertNumberOfCalls(t, "ReadAll", 1)
	repoMock.AssertNumberOfCalls(t, "ReadByID", 0)
	repoMock.AssertNumberOfCalls(t, "Insert", 0)
}
