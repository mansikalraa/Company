package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Database connection
	db, err := gorm.Open("postgres", "user=postgres password=belikemee dbname=try port=5432 sslmode=disable")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database := db.DB()

	err = database.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connection to PostgreSQL was successful!")

	db.AutoMigrate(&Trainee{})
	db.AutoMigrate(&Trainer{})

	DB = db
}
