package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Posts    []Post
}

type Post struct {
	gorm.Model
	UserID      uint   `json:"userId"`
	Description string `json:"description"`
}
