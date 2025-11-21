package rpcserver

import (
	"context"
	"log"
	"net/http"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Config struct {
	Addr string
}

type Server struct {
	httpServer *http.Server
	logger     *log.Logger
}

func New(cfg Config, services Services, logger *log.Logger) *Server {
	if cfg.Addr == "" {
		cfg.Addr = ":8080"
	}
	if logger == nil {
		logger = log.Default()
	}

	handler := buildHandler(services, logger)
	h2cHandler := h2c.NewHandler(handler, &http2.Server{})

	return &Server{
		httpServer: &http.Server{
			Addr:              cfg.Addr,
			Handler:           h2cHandler,
			ReadHeaderTimeout: 10 * time.Second,
		},
		logger: logger,
	}
}

func (s *Server) ListenAndServe() error {
	s.logger.Printf("rpc server listening on %s", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.logger.Printf("rpc server shutting down")
	return s.httpServer.Shutdown(ctx)
}
