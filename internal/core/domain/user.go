package domain

import (
	"time"

	"github.com/google/uuid"
)

// User is an entity that represents a user
type User struct {
	ID        uuid.UUID
	FirstName string
	LastName string
	Email     string
	Password  string
	PhoneNumber string
	Address *Address
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateUserParams struct {
	FirstName     string
	LastName string
	Email    string
	Password string
	PhoneNumber string
	Address 	string
	City      string
	Province  string
	Zip 			string
	Country   string
}