package domain

import (
	"context"
	"time"
)

type User struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	LastModifiedAt time.Time `json:"last_modified_at"`
	CreatedAt      time.Time `json:"created_at"`
}

type UserRepository interface {
	FindById(ctx context.Context, id string) (*User, error)
	FindAll() ([]*User, error)
}

type UserService interface {
	FindById(ctx context.Context, id string) (*User, error)
	FindAll() ([]*User, error)
}
