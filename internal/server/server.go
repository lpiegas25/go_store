package server

import (
	"github.com/go-chi/chi"
	v1 "github.com/lpiegas25/go_store/internal/server/v1"
	"log"
	"net/http"
	"time"
)

// Server is a base server configuration.
type Server struct {
	server *http.Server
}

// New inicialize a new server with configuration.
func New(port string) (*Server, error) {
	r := chi.NewRouter()

	/*
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	 */

	// API routes version 1.
	r.Mount("/api/v1", v1.New())

	serv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := Server{server: serv}

	return &server, nil
}

func (serv *Server) Close() error {
	// TODO: add resource closure.
	return nil
}

// Start the server.
func (serv *Server) Start() {
	log.Printf("Server running on http://localhost%s", serv.server.Addr)
	err := serv.server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}