package domain

import (
	"context"
	"database/sql"
	"time"
)

type User struct {
	ID             string       `json:"id"`
	Name           string       `json:"name"`
	Email          string       `json:"email"`
	LastModifiedAt sql.NullTime `json:"last_modified_at"`
	CreatedAt      time.Time    `json:"created_at"`
}

type UserRepository interface {
	FindById(ctx context.Context, id int64) (*User, error)
	FindAll() ([]*User, error)
	Create() error
	Update() error
	Delete() error
}

type UserService interface {
	FindById(ctx context.Context, id int64) (*User, error)
	FindAll() ([]*User, error)
}
