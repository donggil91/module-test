package service

import (
	"github.com/dongil91/module-test/domain"
)

type userService struct {
	userRepository domain.SellerRepository
}

func NewUserService(u domain.SellerRepository) domain.SellerService {
	return &userService{
		userRepository: u,
	}
}

func (u *userService) FindById(id int64) (*domain.Seller, error) {
	user, err := u.userRepository.FindById(id)
	if err != nil {
		panic(err)
	}
	return user, nil
}

func (u *userService) FindAll() ([]*domain.Seller, error) {
	users, err := u.userRepository.FindAll()
	if err != nil {
		panic(err)
	}
	return users, nil
}

func (u *userService) Create(name string, email string) error {
	return u.userRepository.Create(name, email)
}

func (u *userService) Update(name string, id int64) error {
	return u.userRepository.Update(name, id)
}

func (u *userService) Delete(id int) error {
	return u.userRepository.Delete(id)
}
