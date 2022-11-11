package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Comment string `json:"comment" example:"class"`
	GistID  int    `json:"gist_id" example:"111"`
	Gist    Gist
	UserID  int `json:"user_id" example:"21"`
	User    User
}
