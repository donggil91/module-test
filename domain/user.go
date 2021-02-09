package domain

import (
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

type Reader interface {
	FindById(id int64) (*User, error)
	FindAll() ([]*User, error)
}

type Writer interface {
	Create(name string, email string) error
	Update(name string, email string, id int64) error
	Delete(id int64) error
}

type UserRepository interface {
	Reader
	Writer
}

type UserService interface {
	FindById(id int64) (*User, error)
	FindAll() ([]*User, error)
	Create(name string, email string) error
	Update(name string, email string, id int64) error
	Delete(id int64) error
}
