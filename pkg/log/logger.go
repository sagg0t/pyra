package log

import (
	"context"
	"log/slog"
)

type Level slog.Level

const (
	LevelDebug Level = Level(slog.LevelDebug) // -4
	LevelTrace Level = -2
	LevelInfo  Level = Level(slog.LevelInfo)  // 0
	LevelWarn  Level = Level(slog.LevelWarn)  // 4
	LevelError Level = Level(slog.LevelError) // 8
)

func init() {
	slog.SetLogLoggerLevel(slog.Level(LevelTrace))
}

type Logger struct {
	slogger *slog.Logger
}

func NewLogger() *Logger {
	return &Logger{
		slogger: slog.Default(),
		// slogger: slog.New(slog.NewTextHandler(os.Stderr, nil)),
	}
}

func (l *Logger) Log(ctx context.Context, level Level, msg string, args ...any) {
	l.slogger.Log(ctx, slog.Level(level), msg, args...)
}

func (l *Logger) Debug(msg string, args ...any) {
	l.slogger.Debug(msg, args...)
}

func (l *Logger) DebugContext(ctx context.Context, msg string, args ...any) {
	l.slogger.DebugContext(ctx, msg, args...)
}

func (l *Logger) Trace(msg string, args ...any) {
	l.slogger.Log(context.Background(), slog.Level(LevelTrace), msg, args...)
}

func (l *Logger) TraceContext(ctx context.Context, msg string, args ...any) {
	l.slogger.Log(ctx, slog.Level(LevelTrace), msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.slogger.Error(msg, args...)
}

func (l *Logger) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.slogger.ErrorContext(ctx, msg, args...)
}

func (l *Logger) Info(msg string, args ...any) {
	l.slogger.Info(msg, args...)
}

func (l *Logger) InfoContext(ctx context.Context, msg string, args ...any) {
	l.slogger.InfoContext(ctx, msg, args...)
}

func (l *Logger) Warn(msg string, args ...any) {
	l.slogger.Warn(msg, args...)
}

func (l *Logger) WarnContext(ctx context.Context, msg string, args ...any) {
	l.slogger.WarnContext(ctx, msg, args...)
}

func (l *Logger) LogAttrs(ctx context.Context, level Level, msg string, attrs ...slog.Attr) {
	l.slogger.LogAttrs(ctx, slog.Level(level), msg, attrs...)
}

func (l *Logger) Enabled(ctx context.Context, level Level) bool {
	return l.slogger.Enabled(ctx, slog.Level(level))
}

func (l *Logger) Handler() slog.Handler {
	return l.slogger.Handler()
}

func (l *Logger) With(args ...any) *Logger {
	return &Logger{
		slogger: l.slogger.With(args...),
	}
}

func (l *Logger) WithGroup(name string) *Logger {
	return &Logger{
		slogger: l.slogger.WithGroup(name),
	}
}
