package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `json:"email" example:"lol@kek.com"`
}
