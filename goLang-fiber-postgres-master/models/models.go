package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
}
