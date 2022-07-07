package main

import (
	"net/http"

	"github.com/aneshas/myapp"
)

// NewServer instantiates new server
func NewServer(provisionAccount myapp.ProvisionAccountFunc) *Server {
	return &Server{
		provisionAccount: provisionAccount,
	}
}

// Server represents http server
type Server struct {
	provisionAccount myapp.ProvisionAccountFunc
}

// ProvisionAccount http handler func
func (s *Server) ProvisionAccount(w http.ResponseWriter, r *http.Request) {
	// Delegate to s.provisionAccount(....)
}
