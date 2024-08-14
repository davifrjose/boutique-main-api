package repository

import (
	"context"

	"github.com/davifrjose/boutique-main-api/internal/adpater/storage/postgres"
	"github.com/davifrjose/boutique-main-api/internal/core/domain"
	"github.com/jackc/pgx/v5"
)

type AddressRepository struct {
	db *postgres.DB
}

func NewAddressRepository(db *postgres.DB) *AddressRepository {
	return &AddressRepository{
		db,
	}
}

func scanAddress(row pgx.Row) (*domain.Address, error) {
	var address domain.Address
	err := row.Scan(
		&address.ID,
    &address.UserId,
    &address.Address,
    &address.City,
    &address.Province,
    &address.Zip,
    &address.Country,
    &address.CreatedAt,
    &address.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func (ar *AddressRepository) CreateAddress(ctx context.Context, address *domain.Address)  error {
	query := `
	INSERT INTO addresses
	(id, user_id, address, city, province, zip, country, created_at, updated_at)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	`

	_, err := ar.db.Exec(ctx,query,
		address.ID,
    address.UserId,
    address.Address,
    address.City,
    address.Province,
    address.Zip,
    address.Country,
    address.CreatedAt,
    address.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}