package main

import (
	"log/slog"
	"os"
)

func down() {
	var count uint64 = 0
	if len(os.Args) > 0 {
		count = parseCount(os.Args[0])
	}

	migs, err := migrEngine.Down(count)
	if err != nil {
		slog.Error("error while rolling back migrations", "errors", err)
	}

	for _, m := range migs {
		slog.Info("successfully rolled back a migration", "version", m.Version)
	}
}
