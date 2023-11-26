package user

import (
	"net/http"

	"github.com/labstack/echo/v4"
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

	newUser, err := models.NewUser(requestBody.Username, requestBody.Password, requestBody.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	// err := db.Database.Create(&newUser).Error
	// if err != nil {
	// 	fmt.Println("failed to create User:", err)
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to register user"})
	// }

	return c.JSON(http.StatusCreated, newUser)
}
