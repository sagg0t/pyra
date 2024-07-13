package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"

	"pyra/pkg/log"
)

type traceCtxKey int

const (
	traceStart traceCtxKey = iota
	queryData
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
	ctx = context.WithValue(ctx, queryData, data)

	return ctx
}

func (t *QueryTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	now := time.Now()
	took := now.Sub(ctx.Value(traceStart).(time.Time))

	startData := ctx.Value(queryData).(pgx.TraceQueryStartData)

	t.log.TraceContext(
		ctx,
		"DB query",
		"query", startData.SQL,
		"args", startData.Args,
		"took", took,
		"error", data.Err,
	)
}
