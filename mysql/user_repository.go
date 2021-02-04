package mysql

import (
	"context"
	"database/sql"

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
	return &domain.User{}, nil
}

func (u *userRepository) FindAll() ([]*domain.User, error) {
	return nil, nil
}
