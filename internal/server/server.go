package server

import (
	"context"
	"net/http"
	"time"

	"github.com/JudgeSchnoopy/go-sonar/sonar"
	"github.com/gorilla/mux"
)

// Server serves http responses
type Server struct {
	http     *http.Server
	Registry sonar.Registry
}

// New generates a new server
func New() (Server, error) {
	server := Server{
		http: &http.Server{
			Addr:         ":8080",
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
		Registry: sonar.NewRegistry(),
	}

	server.http.Handler = server.router()

	return server, nil
}

// Start begins the listening service.
func (server *Server) Start() error {
	err := server.http.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

// Stop shuts down the listening service
func (server *Server) Stop(ctx context.Context) {
	server.http.Shutdown(ctx)
}

func (server *Server) router() *mux.Router {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.HandleFunc("/docs", docsHandler).Methods("GET")
	r.HandleFunc("/registry", server.showRegistryHandler).Methods("GET")
	r.HandleFunc("/register", server.registerHandler).Methods("POST")

	return r
}