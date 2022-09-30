package repositories

import (
	"fmt"

	"github.com/WilliamKSilva/unit-tests-go/application/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entities.User) (*entities.User, error)
	Find(id uint) (*entities.User, error)
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

func (repo UserRepositoryDb) Find(id uint) (*entities.User, error) {

	var user entities.User
	repo.Db.Find(&user, "id = ?", id)

	if user.Name != "" {
		return nil, fmt.Errorf("User does not exist!")
	}

	return &user, nil
}
