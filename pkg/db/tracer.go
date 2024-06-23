package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/olehvolynets/pyra/pkg/log"
)

type traceCtxKey int

const (
	traceStart traceCtxKey = iota
)

type QueryTracer struct {
	log *log.Logger
}

func NewQueryTracer(l *log.Logger) *QueryTracer {
	return &QueryTracer{log: l}
}

func (t *QueryTracer) TraceQueryStart(
	ctx context.Context,
	conn *pgx.Conn,
	data pgx.TraceQueryStartData,
) context.Context {
	ctx = context.WithValue(ctx, traceStart, time.Now())

	t.log.Debug("trace start", "query", data.SQL, "args", data.Args)
	return ctx
}

func (t *QueryTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	now := time.Now()
	took := now.Sub(ctx.Value(traceStart).(time.Time))
	t.log.Debug("trace end", "took", took, "error", data.Err)
}
