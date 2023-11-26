package models

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func generateToken(userID uuid.UUID) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = userID.String()

	// Sign the token with a secret key
	signedToken, err := token.SignedString([]byte("your-secret-key")) // Replace with your actual secret key
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func NewUser(username, password, email string) (*User, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	token, err := generateToken(id)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:        id,
		Username:  username,
		Password:  password,
		Email:     email,
		Token:     token,
		Status:    true,
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Username  string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	Email     string    `gorm:"not null;unique"`
	Token     string
	Status    bool
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
