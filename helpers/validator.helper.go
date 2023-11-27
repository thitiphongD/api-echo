package helpers

import (
	"errors"
	"unicode"
)

func StrongPassword(password string) error {
	var (
		hasUppercase bool
		hasLowercase bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUppercase = true
		case unicode.IsLower(char):
			hasLowercase = true
		}
	}

	if !hasUppercase {
		return errors.New("password must contain at least one uppercase letter")
	}

	if !hasLowercase {
		return errors.New("password must contain at least one lowercase letter")
	}

	if len(password) <= 7 {
		return errors.New("password must be longer than 8 characters")
	}

	return nil
}
