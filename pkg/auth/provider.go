package auth

import (
	"context"
	"time"
)

type ProviderKey string

const (
	ProviderGoogleOAuth2 ProviderKey = "google_oauth2"
)

type Provider struct {
	ID     uint64
	UserID uint64

	Name string
	UID  string

	CreatedAt time.Time
	UpdatedAt time.Time
}

type AuthProvidersRepository interface {
	Find(ctx context.Context, name, uid string) (Provider, error)
	Create(ctx context.Context, userId uint64, name, uid string) (uint64, error)
}

type GoogleUser struct {
	UID           string `json:"id"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"verified_email"`
	FirstName     string `json:"given_name"`
	LastName      string `json:"family_name"`
}
