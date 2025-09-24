package services

import (
	"context"
	"errors"

	"github.com/aryadhira/go-fiber-template/internal/interfaces"
	"github.com/aryadhira/go-fiber-template/internal/models"
)

type userService struct {
	userRepo interfaces.UserRepository
}

func NewUserService(userRepo interfaces.UserRepository) interfaces.UserService {
	return &userService{userRepo: userRepo}
}

func (u *userService) GetProfile(ctx context.Context, userID string) (*models.User, error) {
	user, err := u.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
