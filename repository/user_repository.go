package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/dongil91/module-test/domain"
)

type UserRepository struct {
	DB *sql.DB
}

func NewMysqlUserRepository(db *sql.DB) domain.UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) FindById(ctx context.Context, id string) (*domain.User, error) {
	return &domain.User{
		ID:             "1",
		Name:           "Dong Gil",
		Email:          "ndgndg91@gmail.com",
		LastModifiedAt: time.Time{},
		CreatedAt:      time.Time{},
	}, nil
}

func (u *UserRepository) FindAll() ([]*domain.User, error) {
	return nil, nil
}
