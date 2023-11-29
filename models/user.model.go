package models

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/thitiphongD/api-echo/helpers"
	"golang.org/x/crypto/bcrypt"
)

func NewUser(email, username, password string) (*User, error) {
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
		Email:     email,
		Username:  username,
		Password:  string(hashedPassword),
		Token:     token,
		Status:    true,
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

type User struct {
	ID        uuid.UUID `gorm:"primary_key"`
	Username  string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	Email     string    `gorm:"not null;unique"`
	Token     string
	Status    bool
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
