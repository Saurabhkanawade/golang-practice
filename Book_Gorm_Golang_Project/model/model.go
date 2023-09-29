package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title    string `json:"title"`
	Rating   string `json:"rating"`
	AuthorId uint   `gorm:"foreignKey:authorId" json:"-"`
	Author   Author `json:"author"`
}

type Author struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
}
