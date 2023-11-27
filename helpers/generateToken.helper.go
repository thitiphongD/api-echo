package helpers

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func GenerateToken(userID uuid.UUID) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userID.String()

	signedToken, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
