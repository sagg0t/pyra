package log

import (
	"context"
	"encoding/json"
	"log/slog"
	"os"
)

type Level slog.Level

const (
	LevelDebug Level = Level(slog.LevelDebug)
	LevelTrace Level = -2
	LevelInfo  Level = Level(slog.LevelInfo)
	LevelWarn  Level = Level(slog.LevelWarn)
	LevelError Level = Level(slog.LevelError)
)

type Logger struct {
	slogger *slog.Logger
}

func NewLogger() *Logger {
	return &Logger{
		slogger: slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.Level(LevelDebug),
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.TimeKey {
					return slog.Attr{Key: "ts", Value: a.Value}
				}

				if a.Key == slog.LevelKey {
					attr := slog.Attr{Key: "severity"}

					level := a.Value.Any().(slog.Level)
					switch {
					case level == slog.Level(LevelTrace):
						attr.Value = slog.StringValue("TRACE")
					default:
						attr.Value = slog.StringValue(level.String())
					}

					return attr
				}

				return a
			},
		})),
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

func (l *Logger) Inspect(v any) {
	l.InspectContext(context.Background(), v)
}

func (l *Logger) InspectContext(ctx context.Context, v any) {
	prettyValue, err := json.MarshalIndent(v, "", "	")
	if err != nil {
		l.ErrorContext(ctx, "failed to marshal value for inspection", "error", err)
		return
	}

	l.DebugContext(ctx, string(prettyValue))
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
