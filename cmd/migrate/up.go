package main

import (
	"log/slog"
	"os"
)

func up() {
	var count uint64 = 0
	if len(os.Args) > 0 {
		count = parseCount(os.Args[0])
	}

	migs, err := migrEngine.Up(count)
	if err != nil {
		slog.Error("error while applting migrations", "error", err)
	}

	for _, m := range migs {
		slog.Info("successfully applied a migration", "version", m.Version)
	}
}
