package service

import (
	"context"

	"github.com/davifrjose/boutique-main-api/internal/core/domain"
	"github.com/davifrjose/boutique-main-api/internal/core/port"
)

type AddressService struct {
	repo  port.AddressRepository
}

func NewAddressService(repo port.AddressRepository, ) *AddressService {
	return &AddressService{
		repo,
	}
}

func (us *AddressService) Register(ctx context.Context, address *domain.Address) (*domain.Address, error) {
	 err := us.repo.CreateAddress(ctx,address)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}
	return address, err
}