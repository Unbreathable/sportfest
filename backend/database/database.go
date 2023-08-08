package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func Setup() {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	var err error
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("could not connect to the database")
	}

	log.Println("Connection opened to PostgreSQL")

	DBConn.AutoMigrate()

}
