package repository

import (
	"context"

	"github.com/aryadhira/go-fiber-template/internal/interfaces"
	"github.com/aryadhira/go-fiber-template/internal/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Create(ctx context.Context, user *models.User) error {
	return u.db.WithContext(ctx).Create(user).Error
}

func (u *userRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	var user models.User
	err := u.db.WithContext(ctx).Where("id = ?", id).First(&user).Error
	return &user, err
}

func (u *userRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := u.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	return &user, err
}
