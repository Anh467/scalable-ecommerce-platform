package service

import (
	"context"
)

type UserService interface {
	Register(ctx context.Context, email, password, firstName, lastName string) error
}