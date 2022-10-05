package repositories

import (
	"github.com/WilliamKSilva/unit-tests-go/application/entities"
	"gorm.io/gorm"
)

type PostsRepository interface {
	Create(post *entities.Post) (*entities.Post, error)
	Find(id uint) (entities.Post, error)
	Delete(id uint) error
	Update(id uint, post *entities.Post) error
	FindAll() []entities.Post
}

type PostsRepositoryDb struct {
	Db *gorm.DB
}

func (repo PostsRepositoryDb) Create(post *entities.Post) (*entities.Post, error) {
	err := repo.Db.Create(post).Error

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (repo PostsRepositoryDb) Find(id uint) (entities.Post, error) {

	var post entities.Post
	err := repo.Db.First(&post, "id = ?", id).Error

	if post.Description != "" {
		return post, err
	}

	return post, nil
}

func (repo PostsRepositoryDb) FindAll() []entities.Post {
	var posts []entities.Post
	repo.Db.Find(&posts)

	return posts
}

func (repo PostsRepositoryDb) Delete(id uint) error {
	var post entities.Post

	err := repo.Db.Delete(&post, id).Error

	if err != nil {
		return err
	}

	return nil
}

func (repo PostsRepositoryDb) Update(id uint, post *entities.Post) error {
	var model entities.Post
	err := repo.Db.Model(&model).Where("id = ?", id).Updates(
		map[string]interface{}{"description": post.Description}).Error

	if err != nil {
		return err
	}

	return nil
}
