package entities

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID      uint   `json:"userId"`
	Description string `json:"description"`
}
