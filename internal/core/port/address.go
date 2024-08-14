package port

import (
	"context"

	"github.com/davifrjose/boutique-main-api/internal/core/domain"
)

// AddressRepository is an interface for interacting with address-related data
type AddressRepository interface {
	// CreateAddress inserts a new address into the database
	CreateAddress(ctx context.Context, address *domain.Address)  error
}

// AddressService is an interface for interacting with address-related business logic
type AddressService interface {
	// Register registers a new address
	Register(ctx context.Context, address *domain.Address) (*domain.Address, error)
}
