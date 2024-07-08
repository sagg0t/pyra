package log

import "context"

type ctxKey string

const loggerCtxKey ctxKey = "logger"

func CtxWithLogger(ctx context.Context, l *Logger) context.Context {
	return context.WithValue(ctx, loggerCtxKey, l)
}

func FromContext(ctx context.Context) *Logger {
	return ctx.Value(loggerCtxKey).(*Logger)
}
