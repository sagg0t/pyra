package db

import "context"

type Repository[T any] interface {
	BeginTx(context.Context) (DBTX, error)
	WithTx(DBTX) T
}
