package models

import "gorm.io/gorm"

type File struct {
	gorm.Model
	Name    string `json:"name" example:"class"`
	Content string `json:"content" example:"<?php echo 1;?>"`
	GistID  int    `json:"gist_id" example:"11"`
	Gist    Gist
}
