package main

import (
	"github.com/douglastaylorb/favorites-games-api/database"
	"github.com/douglastaylorb/favorites-games-api/routes"
)

func main() {
	database.ConnectDB()
	routes.HandleRequests()
}
