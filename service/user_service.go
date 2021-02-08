package service

import (
	"context"

	"github.com/dongil91/module-test/domain"
)

type userService struct {
	userRepository domain.UserRepository
}

func NewUserService(u domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: u,
	}
}

func (u *userService) FindById(ctx context.Context, id string) (*domain.User, error) {
	user, err := u.userRepository.FindById(ctx, id)
	if err != nil {
		panic(err)
	}
	return user, nil
}

func (u *userService) FindAll() ([]*domain.User, error) {
	return nil, nil
}
