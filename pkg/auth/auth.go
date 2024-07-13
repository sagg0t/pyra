package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"pyra/pkg/log"
	"pyra/pkg/users"
)

type AuthService struct {
	db           *pgxpool.Pool
	log          *log.Logger
	providerRepo AuthProvidersRepository
	userRepo     users.UserRepository
}

func NewService(
	log *log.Logger,
	db *pgxpool.Pool,
	proproviderRepo AuthProvidersRepository,
	userRepo users.UserRepository,
) *AuthService {
	return &AuthService{
		log:          log,
		db:           db,
		providerRepo: proproviderRepo,
		userRepo:     userRepo,
	}
}

func (svc *AuthService) SignIn(ctx context.Context, guser GoogleUser) (users.User, error) {
	var user users.User
	var err error

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
				svc.log.ErrorContext(ctx, "failed to rollback a transaction", "error", rollbackErr)
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

	// User found, so either sign in, or first sign in using Google.
	// Just need to add the auth provider for consistency.
	_, err = svc.db.Exec(
		ctx,
		"INSERT INTO auth_providers (name, uid) VALUES ($1, $2) ON CONFLICT (name, uid) DO nothing;",
		ProviderGoogleOAuth2, guser.UID,
	)

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
