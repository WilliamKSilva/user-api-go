package usecases_test

import (
	"testing"
	"time"

	"github.com/WilliamKSilva/unit-tests-go/application/entities"
	"github.com/WilliamKSilva/unit-tests-go/application/repositories"
	usecases "github.com/WilliamKSilva/unit-tests-go/application/usecases/user"
)

func Create(t *testing.T) {
	var db []*entities.User
	repo := repositories.UserRepositoryFakeDb{Db: db}
	userUseCase := usecases.UserUseCase{UserRepository: repo}

	var user *entities.User

	user.ID = 1
	user.Name = "teste"
	user.Password = "12345"
	user.Email = "teste@teste.com"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	user_created, err := userUseCase.Create(user)

	if err != nil {
		t.Errorf(err.Error())
	}

	if user_created != user {
		t.Errorf("User created wrong!")
	}

}
