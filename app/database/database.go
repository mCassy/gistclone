package database

import (
	"log"
	"os"

	"gitlab.com/mCassy/gistclone/app/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

// connectDb
func ConnectDb() {

	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")

	db.AutoMigrate(&models.Gist{}, &models.User{}, &models.File{}, &models.Comment{})

	DBConn = db

}
