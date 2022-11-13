package dto

import "gitlab.com/mCassy/gistclone/app/models"

type GistRequestBody struct {
	Name        string    `json:"name" example:"class"`
	Description string    `json:"description" example:"Gist short description"`
	IsPublic    bool      `json:"is_public" example:"true"`
	Files       []FileDTO `json:"files"`
}

func (gist GistRequestBody) ToGistModel() *models.Gist {
	return &models.Gist{Name: gist.Name, Description: gist.Description, IsPublic: gist.IsPublic}
}

type GistRequestParams struct {
	GistID int
}
