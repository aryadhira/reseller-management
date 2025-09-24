package interfaces

import (
	"context"
	"time"

	"github.com/aryadhira/go-fiber-template/internal/models"
)

type AuthService interface {
	Register(ctx context.Context, req *models.RegisterRequest) (*models.AuthResponse, error)
	Login(ctx context.Context, req *models.LoginRequest, jwtSecret string, expiresIn time.Duration) (*models.AuthResponse, error)
}

type UserService interface {
	GetProfile(ctx context.Context, userID string) (*models.User, error)
}
