package main

import (
	"log/slog"
	"os"
)

func redo() {
	migs, err := migrEngine.Down(1)
	if err != nil {
		slog.Error("error while rolling back migrations", "error", err)
		os.Exit(1)
	}

	for _, m := range migs {
		slog.Info("successfully rolled back a migration", "version", m.Version)
	}

	migs, err = migrEngine.Up(1)
	if err != nil {
		slog.Error("error while applting migrations", "error", err)
		os.Exit(1)
	}

	for _, m := range migs {
		slog.Info("successfully applied a migration", "version", m.Version)
	}
}
