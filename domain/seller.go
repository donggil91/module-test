package domain

import (
	"database/sql"
	"time"
)

type Seller struct {
	ID             string       `json:"id"`
	Name           string       `json:"name"`
	Email          string       `json:"email"`
	LastModifiedAt sql.NullTime `json:"last_modified_at"`
	CreatedAt      time.Time    `json:"created_at"`
}

type Reader interface {
	FindById(id int64) (*Seller, error)
	FindAll() ([]*Seller, error)
}

type Writer interface {
	Create(name string, email string) error
	Update(name string, id int64) error
	Delete(id int) error
}

type SellerRepository interface {
	Reader
	Writer
}

type SellerService interface {
	FindById(id int64) (*Seller, error)
	FindAll() ([]*Seller, error)
	Create(name string, email string) error
	Update(name string, id int64) error
	Delete(id int) error
}
