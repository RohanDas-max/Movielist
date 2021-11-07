package main

import (
	"movielist/database"
	"movielist/routes"

	"github.com/joho/godotenv"
)

func main() {
	database.Connection()
	routes := routes.Routes()

	envMap, _ := godotenv.Read(".env")

	routes.Run(envMap["PORT"])
}
