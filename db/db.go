package db

import (
	"book-backend/model"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error when load file .env")
		return nil, err
	}
	dsnString := os.Getenv("DB_CONNECTION_STRING")
	db, err := gorm.Open(mysql.Open(dsnString), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection failed:", err)
		return nil, err
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.User{},
		&model.Book{},
		&model.Comment{},
		&model.Tag{},
	)
}
