package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID
	Email         string
	PasswordHash  string
	FirstName     string
	LastName      string
	Phone         *string
	Role          string
	Status        string
	EmailVerified bool
	LastLoginAt   *time.Time
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time
}