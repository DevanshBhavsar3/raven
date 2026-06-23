package auth

import (
	"context"
	"fmt"

	"github.com/DevanshBhavsar3/raven/internal/config"
	"github.com/clerk/clerk-sdk-go/v2"
	clerkUser "github.com/clerk/clerk-sdk-go/v2/user"
)

type AuthService struct{}

func NewAuthService(cfg *config.ApplicationConfig) *AuthService {
	clerk.SetKey(cfg.Auth.SecretKey)

	return &AuthService{}
}

func (s *AuthService) GetUserEmail(ctx context.Context, userID string) (string, error) {
	user, err := clerkUser.Get(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("failed to get user: %w", err)
	}

	if len(user.EmailAddresses) == 0 {
		return "", fmt.Errorf("user %s has no email addresses", userID)
	}

	for _, email := range user.EmailAddresses {
		if user.PrimaryEmailAddressID != nil && email.ID == *user.PrimaryEmailAddressID {
			return email.EmailAddress, nil
		}
	}

	return user.EmailAddresses[0].EmailAddress, nil
}
