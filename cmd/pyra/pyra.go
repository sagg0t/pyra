package main

import (
	"context"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/olehvolynets/pyra/internal/api/foodproducts"
	"github.com/olehvolynets/pyra/internal/server"
)

var port *uint = flag.Uint("port", 42069, "Host the server is running on.")

func init() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
}

func main() {
	flag.Parse()

	appLogger := slog.Default()

	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		done()
		if r := recover(); r != nil {
			appLogger.Error("application panic", "panic", r)
			os.Exit(1)
		}
	}()

	srv, err := server.New(server.WithPort(*port), server.WithLogger(appLogger))
	if err != nil {
		appLogger.Error("failed to create an http server", "error", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/foodProducts", foodproducts.ServerMux())

	staticFs := http.FileServerFS(os.DirFS("./public/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", staticFs))

	err = srv.Start(ctx, server.Logger(mux))
	if err != nil {
		appLogger.Error("failed to start a server", "error", err)
	}
}
