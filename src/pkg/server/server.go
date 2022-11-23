package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	*http.Server
}

func NewServer(port string, handler http.Handler) *Server {
	return &Server{
		&http.Server{
			Addr:           ":" + port,
			Handler:        handler,
			MaxHeaderBytes: 1 << 20, // 1 MB
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	return s.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}
