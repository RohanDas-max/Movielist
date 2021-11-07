package database

import (
	"fmt"
	"log"
	"movielist/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {

	envMap, envErr := godotenv.Read(".env")
	if envErr != nil {
		log.Fatal("error occured on env var", envErr)
	}
	dsn := fmt.Sprintf("%s:%s@tcp%s/%s?charset=utf8mb4&parseTime=True&loc=Local", envMap["USER"], envMap["PASS"], envMap["HOST"], envMap["DATABASE"])

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

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
