package server

import (
	"context"
	"fmt"
	"net/http"
)

// Server holds the global configuration options
type Server struct {
	Address    string
	Port       int
	Version    string
	httpServer *http.Server
}

func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ecommerce-ratings " + s.Version + "\n"))
	}
}

// Routes provides routes to the server
func (s *Server) Routes() *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/", s.handleIndex())
	return router
}

// Start initializes the server
func (s *Server) Start() error {
	address := s.Address
	if address == "" {
		address = "0.0.0.0"
	}
	addr := fmt.Sprintf("%s:%d", address, s.Port)
	s.httpServer = &http.Server{
		Addr:    addr,
		Handler: s.Routes(),
	}
	return s.httpServer.ListenAndServe()
}

// Close terminates the server gracefully
func (s *Server) Close() error {
	return s.httpServer.Shutdown(context.Background())
}
