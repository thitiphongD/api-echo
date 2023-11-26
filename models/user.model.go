package models

import "time"

type User struct {
	ID        uint `gorm:"primarykey"`
	Username  string
	Password  string
	Email     string
	Token     string
	Status    bool
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
