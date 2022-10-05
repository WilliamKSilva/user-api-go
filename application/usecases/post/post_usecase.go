package usecases

import (
	"github.com/WilliamKSilva/unit-tests-go/application/entities"
	"github.com/WilliamKSilva/unit-tests-go/application/repositories"
)

type PostUseCase struct {
	PostRepository repositories.PostsRepository
}

func (u *PostUseCase) Create(post *entities.Post) (*entities.Post, error) {
	post, err := u.PostRepository.Create(post)

	if err != nil {
		return post, err
	}

	return post, nil
}

func (u *PostUseCase) Find(id uint) (entities.Post, error) {
	post, err := u.PostRepository.Find(id)

	if err != nil {
		return post, err
	}

	return post, nil
}

func (u *PostUseCase) FindAll() []entities.Post {
	posts := u.PostRepository.FindAll()

	return posts
}

func (u *PostUseCase) Delete(id uint) error {
	err := u.PostRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func (u *PostUseCase) Update(id uint, post *entities.Post) error {
	err := u.PostRepository.Update(id, post)

	if err != nil {
		return err
	}

	return nil
}
