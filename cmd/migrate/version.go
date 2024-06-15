package main

import "log/slog"

func version() {
	mig, err := migrEngine.CurrentVersion()
	if err != nil {
		slog.Error("failed to retireve the DB migration version", "error", err)
	}

	if mig.Version == "0" {
		slog.Info("no DB migrations have been applied so far", "version", "NONE")
	} else {
		slog.Info("current DB migration version", "version", mig.Version)
	}
}
