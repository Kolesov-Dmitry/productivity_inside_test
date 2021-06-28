package users_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"users.org/internal/users"
)

type responseOneUser struct {
	Error interface{} `json:"error"`
	Data  *users.User `json:"data"`
}

type responseUsers struct {
	Error interface{}   `json:"error"`
	Data  []*users.User `json:"data"`
}

func TestUserController_AddUser(t *testing.T) {
	// Setup UserController
	mux := http.NewServeMux()
	repoMock := makeUserRepositoryMock()
	_ = users.NewUserController(repoMock, mux)

	// Setup Test Server
	srv := httptest.NewServer(mux)
	defer srv.Close()

	// Encode User data
	postBody, _ := json.Marshal(&users.User{
		Name:     "Test",
		LastName: "User",
		Age:      20,
	})

	// Do request
	res, err := http.Post(
		fmt.Sprintf("%s/api/user", srv.URL),
		"application/json",
		bytes.NewBuffer(postBody),
	)

	// Check
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// Read body
	defer res.Body.Close()
	var userResp responseOneUser
	body, _ := ioutil.ReadAll(res.Body)
	_ = json.Unmarshal(body, &userResp)

	// Check body
	assert.Nil(t, userResp.Error)
	assert.NotNil(t, userResp.Data)
	assert.Equal(t, expectedUserID, userResp.Data.ID)
}

func TestUserController_AddUser_WithError(t *testing.T) {
	// Setup UserController
	mux := http.NewServeMux()
	repoMock := makeUserRepositoryMock()
	_ = users.NewUserController(repoMock, mux)

	// Setup Test Server
	srv := httptest.NewServer(mux)
	defer srv.Close()

	// Encode User data wit error
	postBody, _ := json.Marshal(&users.User{
		LastName: "User",
		Age:      20,
	})

	// Do request
	res, err := http.Post(
		fmt.Sprintf("%s/api/user", srv.URL),
		"application/json",
		bytes.NewBuffer(postBody),
	)

	// Check
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// Read body
	defer res.Body.Close()
	var userResp responseOneUser
	body, _ := ioutil.ReadAll(res.Body)
	_ = json.Unmarshal(body, &userResp)

	// Check body
	assert.NotNil(t, userResp.Error)
	assert.Nil(t, userResp.Data)
}

func TestUserController_GetAllUsers(t *testing.T) {
	// Setup UserController
	mux := http.NewServeMux()
	repoMock := makeUserRepositoryMock()
	_ = users.NewUserController(repoMock, mux)

	// Setup Test Server
	srv := httptest.NewServer(mux)
	defer srv.Close()

	// Do request
	res, err := http.Get(fmt.Sprintf("%s/api/user", srv.URL))

	// Check
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// Read body
	defer res.Body.Close()
	var userResp responseUsers
	body, _ := ioutil.ReadAll(res.Body)
	_ = json.Unmarshal(body, &userResp)

	// Check body
	assert.Nil(t, userResp.Error)
	assert.NotNil(t, userResp.Data)
	assert.True(t, reflect.DeepEqual(userResp.Data, expectedUsers))
}

func TestUserController_GetUserByID(t *testing.T) {
	// Setup UserController
	mux := http.NewServeMux()
	repoMock := makeUserRepositoryMock()
	_ = users.NewUserController(repoMock, mux)

	// Setup Test Server
	srv := httptest.NewServer(mux)
	defer srv.Close()

	// Do request
	res, err := http.Get(fmt.Sprintf("%s/api/user?id=%s", srv.URL, expectedUserID))

	// Check
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// Read body
	defer res.Body.Close()
	var userResp responseOneUser
	body, _ := ioutil.ReadAll(res.Body)
	_ = json.Unmarshal(body, &userResp)

	// Check body
	assert.Nil(t, userResp.Error)
	assert.NotNil(t, userResp.Data)
	assert.Equal(t, expectedUserID, userResp.Data.ID)
}
