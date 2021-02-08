package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/dongil91/module-test/domain"
)

type userRepository struct {
	DB *sql.DB
}

func NewMysqlUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (u *userRepository) FindById(ctx context.Context, id string) (*domain.User, error) {
	return &domain.User{
		ID:             "1",
		Name:           "Dong Gil",
		Email:          "ndgndg91@gmail.com",
		LastModifiedAt: time.Time{},
		CreatedAt:      time.Time{},
	}, nil
}

func (u *userRepository) FindAll() ([]*domain.User, error) {
	return nil, nil
}
