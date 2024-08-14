package service

import (
	"context"
	"log"
	"time"

	"github.com/davifrjose/boutique-main-api/internal/core/domain"
	"github.com/davifrjose/boutique-main-api/internal/core/port"

	"github.com/davifrjose/boutique-main-api/internal/core/util"
	"github.com/google/uuid"
)

/**
 * UserService implements port.UserService interface
 * and provides an access to the user repository
 * and cache service
 */
 type UserService struct {
	userRepo  port.UserRepository
	addressRepo port.AddressRepository
}

// NewUserService creates a new user service instance
func NewUserService(userRepo  port.UserRepository, addressRepo port.AddressRepository) *UserService {
	return &UserService{
		userRepo,
		addressRepo,
	}
}

// Register creates a new user
func (us *UserService) Register(ctx context.Context, params *domain.CreateUserParams) (*domain.User, error) {
	hashedPassword, err := util.HashPassword(params.Password)
	if err != nil {
		return nil, domain.ErrInternal
		
	}

	CreatedAt := time.Now()
	UpdatedAt := time.Now()

	user := &domain.User{
		ID: uuid.New(),
		FirstName: params.FirstName,
		LastName: params.LastName,
		Email: params.Email,
		Password: hashedPassword,
		PhoneNumber: params.PhoneNumber,
		CreatedAt: CreatedAt,
		UpdatedAt: UpdatedAt,
	}
	
	err = us.userRepo.CreateUser(ctx, user)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	address := &domain.Address{
		ID: uuid.New(),
		UserId: user.ID,
		Address: params.Address,
		City: params.City,
		Province: params.Province,
		Zip: params.Zip,
		Country: params.Country,
		CreatedAt: CreatedAt,
		UpdatedAt: UpdatedAt,
	}

	 err = us.addressRepo.CreateAddress(ctx, address)
	 if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		log.Printf("address")
		return nil, domain.ErrInternal
	 }

	 user.Address = address

	return user, nil
}

// GetUser gets a user by ID
func (us *UserService) GetUser(ctx context.Context, id uuid.UUID) (*domain.User, error) {

	user, err := us.userRepo.GetUserByID(ctx, id)
	if err != nil {
		if err == domain.ErrDataNotFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	return user, nil
}