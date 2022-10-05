package repositories

import "github.com/WilliamKSilva/unit-tests-go/application/entities"

type UserRepositoryFakeDb struct {
	Db []*entities.User
}

func (repo UserRepositoryFakeDb) Create(user *entities.User) (*entities.User, error) {
	user_array := append(repo.Db, user)

	user_created := user_array[len(user_array)-1]

	return user_created, nil
}

func (repo UserRepositoryFakeDb) Find(id uint) (entities.User, error) {
	users := repo.Db
	var user entities.User
	for i := 0; i < len(users); i++ {
		if users[i].ID == id {
			return *users[i], nil
		}
	}

	return user, nil
}

func (repo UserRepositoryFakeDb) FindAll() []entities.User {
	var users []entities.User
	repo.Db.Find(&users)

	return users
}

func (repo UserRepositoryFakeDb) Delete(id uint) error {
	var user entities.User

	err := repo.Db.Delete(&user, id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo UserRepositoryFakeDb) Update(id uint, user *entities.User) error {
	var model entities.User
	err := repo.Db.Model(&model).Where("id = ?", id).Updates(
		map[string]interface{}{"name": user.Name, "email": user.Email, "password": user.Password}).Error

	if err != nil {
		return err
	}

	return nil
}
