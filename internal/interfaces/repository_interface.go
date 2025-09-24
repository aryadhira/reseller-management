package interfaces

import (
	"context"

	"github.com/aryadhira/go-fiber-template/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByID(ctx context.Context, id string) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
}
