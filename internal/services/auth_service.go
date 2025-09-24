package services

import (
	"context"
	"errors"
	"time"

	"github.com/aryadhira/go-fiber-template/internal/interfaces"
	"github.com/aryadhira/go-fiber-template/internal/models"
	"github.com/aryadhira/go-fiber-template/internal/utils"
	"github.com/google/uuid"
)

type authService struct {
	userRepo interfaces.UserRepository
}

func NewAuthService(userRepo interfaces.UserRepository) interfaces.AuthService {
	return &authService{userRepo: userRepo}
}

func (a *authService) Register(ctx context.Context, req *models.RegisterRequest) (*models.AuthResponse, error) {
	// Check if user already exists
	existing, err := a.userRepo.FindByEmail(ctx, req.Email)
	if err == nil && existing != nil {
		return nil, errors.New("user already exists with this email")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Create user
	user := &models.User{
		ID:       uuid.NewString(),
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := a.userRepo.Create(ctx, user); err != nil {
		return nil, errors.New("failed to create user")
	}

	return &models.AuthResponse{User: user}, nil
}
func (a *authService) Login(ctx context.Context, req *models.LoginRequest, jwtSecret string, expiresIn time.Duration) (*models.AuthResponse, error) {
	// Find user by email
	user, err := a.userRepo.FindByEmail(ctx, req.Email)
	if err != nil || user == nil {
		return nil, errors.New("invalid credentials")
	}

	// Check password
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user, jwtSecret, expiresIn)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &models.AuthResponse{
		User:        user,
		AccessToken: token,
	}, nil
}
