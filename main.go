package main

import (
	"github.com/douglastaylorb/favorites-games-api/database"
	"github.com/douglastaylorb/favorites-games-api/routes"
)

// @title Favorite Games API
// @version 1.0
// @description This is a sample server for a favorite games API.
// @host localhost:8080
// @BasePath /

func main() {
	database.ConnectDB()
	routes.HandleRequests()
}
