package handlers

import (
	"fmt"
	"gistclone/app/database"
	"gistclone/app/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
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
// @Success 200 {object} ResponseHTTP{data=[]models.Gist}
// @Failure 404 {object} ResponseHTTP{}
// @Failure 503 {object} ResponseHTTP{}
// @Router /v1/gists/{id} [get]
func GetGistByID(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	gist := new(models.Gist)
	if err := db.First(&gist, id).Error; err != nil {
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

	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Success get gist by ID.",
		Data:    *gist,
	})
}
