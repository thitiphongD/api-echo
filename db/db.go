package db

import (
	"fmt"
	"os"

	"github.com/thitiphongD/api-echo/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func InitDB() (*gorm.DB, error) {
	dbHost := os.Getenv("POSTGRES_HOST")
	dbName := os.Getenv("POSTGRES_DB")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", dbHost, dbUser, dbName, dbPort, dbPassword)

	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	Database = db

	return db, nil
}
