package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"time"

	pyralog "github.com/olehvolynets/pyra/pkg/log"
)

type Server struct {
	ip       string
	port     uint
	listener net.Listener
	log      *slog.Logger
}

type ServerOption func(*Server)

func New(opts ...ServerOption) (*Server, error) {
	// FIX: craete a real error logger
	s := &Server{port: 3000, log: slog.Default()}

	for _, opt := range opts {
		opt(s)
	}

	addr := fmt.Sprintf("localhost:%d", s.port)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("failed to create a listener on %s: %w", addr, err)
	}

	s.ip = listener.Addr().(*net.TCPAddr).IP.String()
	s.listener = listener

	return s, nil
}

func WithPort(port uint) ServerOption {
	return func(s *Server) {
		s.port = port
	}
}

func WithLogger(l *slog.Logger) ServerOption {
	return func(s *Server) {
		s.log = l
	}
}

func (s *Server) Start(ctx context.Context, handler http.Handler) error {
	// TODO: configure other options like TLS, remaining timeouts, etc.
	httpServer := http.Server{
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       2 * time.Second,
		// FIX: pass a real error logger
		ErrorLog: log.Default(),
		Handler:  handler,
	}

	errCh := make(chan error, 1)

	s.log.Info("", pyralog.Key, pyralog.ServerStartEvent, "ip", s.ip, "port", s.port)

	go func() {
		<-ctx.Done()

		s.log.Debug("context closed")

		shutdownCtx, done := context.WithTimeout(context.Background(), 5*time.Second)
		defer done()

		s.log.Info("shutting down...", pyralog.Key, pyralog.ServerShutdownEvent)

		errCh <- httpServer.Shutdown(shutdownCtx)

		s.log.Info("stopped", pyralog.Key, pyralog.ServerStopEvent)
	}()

	if err := httpServer.Serve(s.listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}
