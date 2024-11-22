package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

type Config struct {
	Address    string `toml:"address"`
	TLSEnabled bool   `toml:"tls_enabled"`
	CertFile   string `toml:"cert"`
	KeyFile    string `toml:"key"`
}

func getServer(cfg Config, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:           cfg.Address,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
}

func (s *Server) Run(cfg Config, handler http.Handler) error {
	s.httpServer = getServer(cfg, handler)

	if cfg.TLSEnabled {
		return s.httpServer.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile)
	} else {
		return s.httpServer.ListenAndServe()
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
