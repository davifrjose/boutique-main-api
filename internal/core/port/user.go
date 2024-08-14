package port

import (
	"context"

	"github.com/davifrjose/boutique-main-api/internal/core/domain"
	"github.com/google/uuid"
)

// UserRepository is an interface for interacting with user-related data
type UserRepository interface {
	// CreateUser inserts a new user into the database
	CreateUser(ctx context.Context, user *domain.User)  error
	// GetUserByID selects a user by id
	GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	// GetUserByEmail selects a user by email
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
}

// UserService is an interface for interacting with user-related business logic
type UserService interface {
	// Register registers a new user
	Register(ctx context.Context, user *domain.CreateUserParams) (*domain.User, error)
	// GetUser returns a user by id
	GetUser(ctx context.Context, id uuid.UUID) (*domain.User, error)
}

