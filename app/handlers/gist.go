package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/mCassy/gistclone/app/database"
	dto "gitlab.com/mCassy/gistclone/app/dto/gist"
	"gitlab.com/mCassy/gistclone/app/models"
)

type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

// GetGistByID is a function to get a gist by ID
// @Summary Get gist by ID
// @Description Get gist by ID
// @Tags gists
// @Accept json
// @Produce json
// @Param id path int true "Gist ID"
// @Success 200 {object} ResponseHTTP{data=[]dto.GistResponse}
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /api/v1/gists/{id} [get]
func GetGistByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	gist := new(models.Gist)
	if err := db.Preload("Files").First(&gist, id).Error; err != nil {
		switch err.Error() {
		case "record not found":
			return c.Status(http.StatusNotFound).JSON(ResponseHTTP{
				Success: false,
				Message: fmt.Sprintf("Gist with ID %v not found.", id),
				Data:    nil,
			})
		default:
			return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
				Success: false,
				Message: err.Error(),
				Data:    nil,
			})

		}
	}

	GistResponse := dto.CreateGistResponse(*gist)

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success get gist by ID.",
		Data:    GistResponse,
	})
}

// CreateGist creates a new gist data
// @Summary Register a new gist
// @Description Register gist
// @Tags gists
// @Accept json
// @Produce json
// @Param gist body dto.GistRequestBody true "Register gist"
// @Success 200 {object} ResponseHTTP{data=dto.GistResponse}
// @Failure 400 {object} ResponseHTTP{}
// @Router /api/v1/gists [post]
func CreateGist(c *fiber.Ctx) error {
	db := database.DBConn

	gist := new(models.Gist)
	if err := c.BodyParser(&gist); err != nil {
		return c.Status(http.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
	}

	db.Preload("Files").Create(gist)

	GistResponse := dto.CreateGistResponse(*gist)

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success register a gist.",
		Data:    GistResponse,
	})
}

// GetGistList is a function to list of gists
// @Summary Get list of gists
// @Description Get list of public gists
// @Param offset query int true "Offset"
// @Param limit query int true "Limit"
// @Tags gists
// @Accept json
// @Produce json
// @Success 200 {object} ResponseHTTP{data=[]dto.GistResponse}
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /api/v1/gists [get]
func GetGistList(c *fiber.Ctx) error {
	db := database.DBConn

	//TODO: implement pagination
	offset, _ := strconv.Atoi(c.Params("offset"))
	limit, _ := strconv.Atoi(c.Params("limit"))

	var gists models.GistList

	if res := db.Preload("Files").Find(&gists).Limit(limit).Offset(offset); res.Error != nil {
		return c.Status(http.StatusServiceUnavailable).JSON(ResponseHTTP{
			Success: false,
			Message: res.Error.Error(),
			Data:    nil,
		})
	}

	GistResponse := dto.CreateGistListResponse(gists)

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success get gist by ID.",
		Data:    GistResponse,
	})
}
