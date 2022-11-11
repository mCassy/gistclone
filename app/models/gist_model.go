package models

import "gorm.io/gorm"

type Gist struct {
	gorm.Model
	Name        string `json:"name" example:"class"`
	Description string `json:"description" example:"Gist short description"`
	IsPublic    bool   `json:"is_public" example:"true"`
	StarsCnt    int32  `json:"stars_count" example:"21"`
	ForksCnt    int32  `json:"forks_count" example:"21"`
}
