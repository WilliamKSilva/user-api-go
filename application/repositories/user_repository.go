package repositories

import (
	"github.com/WilliamKSilva/unit-tests-go/application/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entities.User) (*entities.User, error)
	Find(id uint) (entities.User, error)
	Delete(id uint) error
	Update(id uint, user *entities.User) error
	FindAll() []entities.User
}

type UserRepositoryDb struct {
	Db *gorm.DB
}

func (repo UserRepositoryDb) Create(user *entities.User) (*entities.User, error) {
	err := repo.Db.Create(user).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo UserRepositoryDb) Find(id uint) (entities.User, error) {

	var user entities.User
	err := repo.Db.First(&user, "id = ?", id).Error

	if user.Name != "" {
		return user, err
	}

	return user, nil
}

func (repo UserRepositoryDb) FindAll() []entities.User {
	var users []entities.User
	repo.Db.Find(&users)

	return users
}

func (repo UserRepositoryDb) Delete(id uint) error {
	var user entities.User

	err := repo.Db.Delete(&user, id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo UserRepositoryDb) Update(id uint, user *entities.User) error {
	var model entities.User
	err := repo.Db.Model(&model).Where("id = ?", id).Updates(
		map[string]interface{}{"name": user.Name, "email": user.Email, "password": user.Password}).Error

	if err != nil {
		return err
	}

	return nil
}
