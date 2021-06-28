package users

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// userResponse holds response to "api/user" request
type userResponse struct {
	Error interface{} `json:"error"`
	Data  interface{} `json:"data"`
}

// UserController takes handle of users requests
type UserController struct {
	userService *UserService
}

// NewUserController creates new UserController instance
func NewUserController(r UserRepository, mux *http.ServeMux) *UserController {
	controller := &UserController{
		userService: NewUserService(r),
	}

	mux.HandleFunc("/api/user", controller.userHandler)

	return controller
}

// userHandler '/api/user' request handler
func (c *UserController) userHandler(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.userGetHandler(rw, r)

	case http.MethodPost:
		c.userPostHandler(rw, r)

	default:
		http.Error(rw, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// userGetHandler '/api/user' GET request handler
func (c *UserController) userGetHandler(rw http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	users, err := c.userService.GetUser(r.Context(), id)
	if err != nil {
		writeUserResponse(rw, err, nil)
		return
	}

	if len(users) == 0 {
		writeUserResponse(rw, nil, nil)
		return
	}

	if id != "" {
		writeUserResponse(rw, nil, users[0])
	} else {
		writeUserResponse(rw, nil, users)
	}
}

// userPostHandler '/api/user' POST request handler
func (c *UserController) userPostHandler(rw http.ResponseWriter, r *http.Request) {
	// read body
	user, err := readUserFromBody(r.Body)
	if err != nil {
		log.Println(err)

		writeUserResponse(rw, "Unable to read request body", nil)
		return
	}
	_ = r.Body.Close()

	// check obligatory fields
	if user.Name == "" {
		writeUserResponse(rw, "Field 'name' is not set", nil)
		return
	}

	if user.Age == 0 {
		writeUserResponse(rw, "Field 'age' is not set", nil)
		return
	}

	// save user
	if err := c.userService.AddUser(r.Context(), user); err != nil {
		writeUserResponse(rw, err, nil)
		return
	}

	// success
	writeUserResponse(rw, nil, user)
}

// writeUserResponse writes response to user request
func writeUserResponse(rw http.ResponseWriter, err interface{}, data interface{}) {
	bytes, _ := json.Marshal(&userResponse{
		Error: err,
		Data:  data,
	})

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")

	if _, err := rw.Write(bytes); err != nil {
		log.Println(err)
	}
}

// reads request body and marshals it to User model
func readUserFromBody(r io.Reader) (*User, error) {
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, err
	}

	return &user, nil
}
