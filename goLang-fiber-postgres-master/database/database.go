package database

import (
	"github.com/mohdaalam005/fiber/helpers"
	"github.com/mohdaalam005/fiber/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect() {
	var url = os.Getenv("url")
	var err error
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	helpers.CheckNilError(err)
	DB.AutoMigrate(&models.Student{})
}
