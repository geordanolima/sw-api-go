package main

import (
	"sw-api-go/database"
	"sw-api-go/routes"
)

func main() {
	database.ConectDB()
	routes.HandleRequest()
}
