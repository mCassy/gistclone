package dto

import (
	"gitlab.com/mCassy/gistclone/app/models"
)

type GistResponse struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" example:"class"`
	Description string    `json:"description" example:"Gist short description"`
	IsPublic    bool      `json:"is_public" example:"true"`
	ForksCnt    int       `json:"forks"`
	StarsCnt    int       `json:"stars"`
	Files       []FileDTO `json:"files"`
}

func CreateGistResponse(gist models.Gist) GistResponse {
	var gistFiles = []FileDTO{}

	for _, file := range gist.Files {
		gistFiles = append(gistFiles, FileDTO{Name: file.Name, Content: file.Content})
	}
	return GistResponse{
		ID:          int(gist.ID),
		Name:        gist.Name,
		Description: gist.Description,
		IsPublic:    gist.IsPublic,
		ForksCnt:    gist.ForksCnt,
		StarsCnt:    gist.StarsCnt,
		Files:       gistFiles,
	}
}

type GistListResponse []*GistResponse

func CreateGistListResponse(gists models.GistList) GistListResponse {
	gistList := GistListResponse{}
	for _, g := range gists {
		gist := CreateGistResponse(*g)
		gistList = append(gistList, &gist)
	}
	return gistList
}
