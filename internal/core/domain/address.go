package domain

import (
	"time"

	"github.com/google/uuid"
)

// Address is an entity that represents a userAddress
type Address struct {
	ID        uuid.UUID
	UserId 		uuid.UUID		
	Address 	string
	City      string
	Province  string
	Zip 			string
	Country   string
	CreatedAt time.Time
	UpdatedAt time.Time
}