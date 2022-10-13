package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Print("No .env file detected trying to access os variabels...")
	}

	if len(os.Getenv("DB_NAME")) == 0 {
		log.Fatal("DB_NAME not set.")
	}

	connection, err := gorm.Open(sqlite.Open(os.Getenv("DB_NAME")), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	DB = connection
}
