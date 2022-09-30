package usecases

import (
	"github.com/WilliamKSilva/unit-tests-go/application/entities"
	"github.com/WilliamKSilva/unit-tests-go/application/repositories"
)

type UserUseCase struct {
	UserRepository repositories.UserRepository
}

func (u *UserUseCase) Create(user *entities.User) (*entities.User, error) {
	user, err := u.UserRepository.Create(user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserUseCase) Find(id uint) (*entities.User, error) {
	user, err := u.UserRepository.Find(id)

	if err != nil {
		return user, err
	}

	return user, nil
}
