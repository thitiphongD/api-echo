package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thitiphongD/api-echo/db"
	"github.com/thitiphongD/api-echo/helpers"
	"github.com/thitiphongD/api-echo/models"
	"github.com/thitiphongD/api-echo/requests"
)

func RegisterUser(c echo.Context) error {
	requestBody := &requests.RequestRegisterUser{}
	if err := c.Bind(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	if requestBody.Password != requestBody.ConfirmPassword {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Password and confirm password do not match"})
	}

	newUser, err := models.CreateNewUser(requestBody.Username, requestBody.Password, requestBody.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	if err := db.Database.Create(&newUser).Error; err != nil {
		if helpers.DuplicateKeyError(err) {
			return c.JSON(http.StatusConflict, map[string]string{"error": "Email already exists"})
		}

		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user in the database"})
	}

	return c.JSON(http.StatusCreated, newUser)
}
