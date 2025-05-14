package auth

import (
	"context"

	"pyra/pkg/db"
)

type ProviderRepository struct {
	db db.DBTX
}

func NewProviderRepository(db db.DBTX) *ProviderRepository {
	return &ProviderRepository{
		db: db,
	}
}

func (svc *ProviderRepository) Find(ctx context.Context, name, uid string) (Provider, error) {
	row := svc.db.QueryRowContext(
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

func (svc *ProviderRepository) Create(ctx context.Context, userId uint64, name, uid string) (uint64, error) {
	row := svc.db.QueryRowContext(
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
