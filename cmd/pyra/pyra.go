package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/olehvolynets/pyra/pkg/api"
	"github.com/olehvolynets/pyra/pkg/db"
	"github.com/olehvolynets/pyra/pkg/log"
	"github.com/olehvolynets/pyra/pkg/server"
)

var port *uint = flag.Uint("port", 42069, "Host the server is running on.")

func main() {
	flag.Parse()

	appLogger := log.NewLogger()

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
		os.Exit(1)
	}

	dbConf := db.NewConfig("postgres")
	dbPool, err := db.CreatePool(ctx, dbConf, appLogger)
	if err != nil {
		appLogger.Error("failed to create a DB pool", "error", err)
		os.Exit(1)
	}

	mux := api.Routes(dbPool, appLogger)

	staticFs := http.FileServerFS(os.DirFS("./public/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", staticFs))

	err = srv.Start(ctx, server.Logger(appLogger, mux))
	if err != nil {
		appLogger.Error("failed to start a server", "error", err)
	}
}
