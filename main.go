package main

import (
	"movielist/database"
	"movielist/routes"
)

func main() {
	database.Connection()
	routes := routes.Routes()

	routes.Run()
}
