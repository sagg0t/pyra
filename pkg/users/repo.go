package users

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (svc *UserRepository) FindById(ctx context.Context, id uint64) (user User, err error) {
	row := svc.db.QueryRow(ctx, "SELECT * FROM users WHERE id = $1 LIMIT 1;", id)

	err = row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return
}

func (svc *UserRepository) FindByEmail(ctx context.Context, email string) (user User, err error) {
	row := svc.db.QueryRow(ctx, "SELECT * FROM users WHERE email = $1 LIMIT 1;", email)

	err = row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return
}

func (svc *UserRepository) Create(ctx context.Context, user User) (uint64, error) {
	row := svc.db.QueryRow(
		ctx,
		"INSERT INTO users (email, first_name, last_name) VALUES ($1, $2, $3) RETURNING id",
		user.Email, user.FirstName, user.LastName,
	)

	var newID uint64
	if err := row.Scan(&newID); err != nil {
		return 0, err
	}

	return newID, nil
}
