package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"pyra/pkg/users"
)

type AuthService struct {
	db           *pgxpool.Pool
	providerRepo ProviderCreator
	userRepo     UserCreateFinder
}

type ProviderCreator interface {
	Create(ctx context.Context, userId uint64, name, uid string) (uint64, error)
}

type (
	UserByEmailFinder interface {
		FindByEmail(ctx context.Context, email string) (users.User, error)
	}
	UserCreator interface {
		Create(ctx context.Context, params users.User) (uint64, error)
	}
	UserCreateFinder interface {
		UserByEmailFinder
		UserCreator
	}
)

func NewService(
	db *pgxpool.Pool,
	proproviderRepo ProviderCreator,
	userRepo UserCreateFinder,
) *AuthService {
	return &AuthService{
		db:           db,
		providerRepo: proproviderRepo,
		userRepo:     userRepo,
	}
}

func (svc *AuthService) SignIn(ctx context.Context, guser GoogleUser) (user users.User, err error) {
	tx, err := svc.db.Begin(ctx)
	if err != nil {
		return user, fmt.Errorf("failed to acquire a transaction - %w", err)
	}
	// Defering commit/rollback. Will rollback if err != nil.
	defer func() {
		if err == nil {
			err = tx.Commit(ctx)
		} else {
			rollbackErr := tx.Rollback(ctx)
			if rollbackErr != nil {
				err = fmt.Errorf("failed to rollback a transaction: %w", rollbackErr)
			}
		}
	}()

	user, err = svc.userRepo.FindByEmail(ctx, guser.Email)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return users.User{}, fmt.Errorf("failed to find user - %w", err)
		}

		// User not found, i.e. sign up flow.
		user, err = svc.CreateUser(ctx, guser)
		if err != nil {
			return users.User{}, err
		}

		_, err := svc.CreateProvider(ctx, user, guser)
		if err != nil {
			return user, fmt.Errorf("failed to create provider - %w", err)
		}

		return user, nil
	}

	return user, err
}

func (svc *AuthService) CreateUser(ctx context.Context, guser GoogleUser) (users.User, error) {
	user := users.User{
		Email:     guser.Email,
		FirstName: guser.FirstName,
		LastName:  guser.LastName,
	}

	id, err := svc.userRepo.Create(ctx, user)
	if err != nil {
		return users.User{}, err
	}

	user.ID = id

	return user, nil
}

func (svc *AuthService) CreateProvider(ctx context.Context, user users.User, guser GoogleUser) (Provider, error) {
	id, err := svc.providerRepo.Create(ctx, user.ID, string(ProviderGoogleOAuth2), guser.UID)
	if err != nil {
		return Provider{}, err
	}

	return Provider{
		ID:     id,
		UserID: user.ID,
		Name:   string(ProviderGoogleOAuth2),
		UID:    guser.UID,
	}, nil
}
