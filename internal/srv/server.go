package srv

import (
	"context"
	"fmt"
	"net/http"

	"users.org/internal/users"
)

// Server is a HTTP server
type Server struct {
	server          *http.Server
	usersController *users.UserController
}

// NewServer creates new Server instance
func NewServer(port int) *Server {
	mux := http.NewServeMux()
	usersController := users.NewUserController(
		users.NewUserRepositoryMem(),
		mux,
	)

	return &Server{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: mux,
		},
		usersController: usersController,
	}
}

// Start starts HTTP server
func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

// Stop closes HTTP server
func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
