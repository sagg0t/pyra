package auth

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type providerRepo struct {
	db *pgxpool.Pool
}

func NewProviderRepository(db *pgxpool.Pool) AuthProvidersRepository {
	return &providerRepo{
		db: db,
	}
}

func (svc *providerRepo) Find(ctx context.Context, name, uid string) (Provider, error) {
	row := svc.db.QueryRow(
		ctx,
		"SELECT * FROM auth_providers WHERE name = $1 AND uid = $2 LIMIT 1",
		name, uid,
	)

	provider := Provider{}

	err := row.Scan(
		&provider.ID,
		&provider.UserID,
		&provider.Name,
		&provider.UID,
		&provider.CreatedAt,
		&provider.UpdatedAt,
	)

	return provider, err
}

func (svc *providerRepo) Create(ctx context.Context, userId uint64, name, uid string) (uint64, error) {
	row := svc.db.QueryRow(
		ctx,
		"INSERT INTO auth_providers (user_id, name, uid) VALUES ($1, $2, $3) RETURNING id",
		userId, name, uid,
	)

	var newID uint64
	if err := row.Scan(&newID); err != nil {
		return 0, err
	}

	return newID, nil
}
