package repository

import (
	"context"
	"log"

	"github.com/davifrjose/boutique-main-api/internal/adpater/storage/postgres"
	"github.com/davifrjose/boutique-main-api/internal/core/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

/**
 * UserRepository implements port.UserRepository interface
 * and provides an access to the postgres database
 */
 type UserRepository struct {
	db *postgres.DB
}

// NewUserRepository creates a new user repository instance
func NewUserRepository(db *postgres.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

// CreateUser creates a new user in the database
func (ur *UserRepository) CreateUser(ctx context.Context, user *domain.User)  error {
	query := ` 
	INSERT INTO users
	(id, first_name, last_name, email, password, phone_number, created_at, updated_at)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
	`

	_,err := ur.db.Exec(ctx, query, 
		user.ID,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.PhoneNumber,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		log.Printf(err.Error())
		return err
	}

	return nil
}

func scanUser(row pgx.Row) (*domain.User, error) {
	var user domain.User
	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.PhoneNumber,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByID gets a user by ID from the database
func (ur *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	query := `
	SELECT * FROM users u WHERE u.id =$1
	`

	row := ur.db.QueryRow(ctx, query, id)
	user, err := scanUser(row)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByEmailAndPassword gets a user by email from the database
func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	query := `
	SELECT * FROM users u WHERE u.email =$1
	`

	row := ur.db.QueryRow(ctx, query, email)
	user, err := scanUser(row)

	if err != nil {
		return nil, err
	}

	return user, nil

}