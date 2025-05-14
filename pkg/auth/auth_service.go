package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"pyra/pkg/db"
)

type AuthService struct {
	db           db.DBTX
	providerRepo ProviderRepository
	userRepo     UserRepository
}

func NewService(
	db db.DBTX,
	proproviderRepo ProviderRepository,
	userRepo UserRepository,
) *AuthService {
	return &AuthService{
		db:           db,
		providerRepo: proproviderRepo,
		userRepo:     userRepo,
	}
}

func (svc *AuthService) SignIn(ctx context.Context, guser GoogleUser) (user User, err error) {
	tx, err := svc.db.BeginTx(ctx, nil)
	if err != nil {
		return user, fmt.Errorf("failed to acquire a transaction - %w", err)
	}
	// Defering commit/rollback. Will rollback if err != nil.
	defer func() {
		if err == nil {
			err = tx.Commit()
		} else {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				err = fmt.Errorf("failed to rollback a transaction: %w", rollbackErr)
			}
		}
	}()

	user, err = svc.userRepo.FindByEmail(ctx, guser.Email)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return User{}, fmt.Errorf("failed to find user - %w", err)
		}

		// User not found, i.e. sign up flow.
		user, err = svc.CreateUser(ctx, guser)
		if err != nil {
			return User{}, err
		}

		_, err := svc.CreateProvider(ctx, user, guser)
		if err != nil {
			return user, fmt.Errorf("failed to create provider - %w", err)
		}

		return user, nil
	}

	return user, err
}

func (svc *AuthService) CreateUser(ctx context.Context, guser GoogleUser) (User, error) {
	user := User{
		Email:     guser.Email,
		FirstName: guser.FirstName,
		LastName:  guser.LastName,
	}

	id, err := svc.userRepo.Create(ctx, user)
	if err != nil {
		return User{}, err
	}

	user.ID = id

	return user, nil
}

func (svc *AuthService) CreateProvider(ctx context.Context, user User, guser GoogleUser) (Provider, error) {
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
