package main

import (
	"log/slog"
	"os"
)

var createFailureMsg = "pyra: failed to create a DB migration"

func createMigration() {
	fileNames, err := migrEngine.CreateMigration(os.Args[0])
	if err != nil {
		slog.Error(createFailureMsg, "error", err)
		os.Exit(1)
	}

	slog.Info("created", "filename", fileNames[0])
	slog.Info("created", "filename", fileNames[1])
}
