package database

import (
	"fmt"
	"log"
	"movielist/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {

	envMap, envErr := godotenv.Read(".env")
	if envErr != nil {
		log.Fatal("error occured on env var", envErr)
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata", envMap["POSTGRESQL_ADDON_HOST"], envMap["POSTGRESQL_ADDON_USER"], envMap["POSTGRESQL_ADDON_PASSWORD"], envMap["POSTGRESQL_ADDON_DB"], envMap["POSTGRESQL_ADDON_PORT"])
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("db con error", err)
	}

	var movies *model.Movielist
	var ratings *model.Ratings
	db.AutoMigrate(&movies)
	db.AutoMigrate(&ratings)

	DB = db

	fmt.Println("DB connected")

}
