package usecases

import (
	"github.com/WilliamKSilva/user-api-go/application/entities"
	"github.com/WilliamKSilva/user-api-go/application/repositories"
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

func (u *UserUseCase) Find(id uint) (entities.User, error) {
	user, err := u.UserRepository.Find(id)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserUseCase) FindAll() []entities.User {
	users := u.UserRepository.FindAll()

	return users
}

func (u *UserUseCase) Delete(id uint) error {
	err := u.UserRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserUseCase) Update(id uint, user *entities.User) error {
	err := u.UserRepository.Update(id, user)

	if err != nil {
		return err
	}

	return nil
}
