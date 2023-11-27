package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/thitiphongD/api-echo/helpers"
	"golang.org/x/crypto/bcrypt"
)

func CreateNewUser(username, password, email string) (*User, error) {
	if err := helpers.StrongPassword(password); err != nil {
		return nil, err
	}

	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	token, err := helpers.GenerateToken(id)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:        id,
		Username:  username,
		Password:  string(hashedPassword),
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
