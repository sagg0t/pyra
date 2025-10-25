package test

import (
	"io"
	"log/slog"
	"os"
	"pyra/pkg/log"
)

func NewLogger(sink io.Writer) *log.Logger {
	h := slog.DiscardHandler

	if os.Getenv("TEST_LOGGING") == "on" {
		h = slog.NewTextHandler(sink, nil)
	}

	return log.NewLoggerFromHandler(h)
}
