package routes

import (
	"github.com/douglastaylorb/favorites-games-api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	// Config cors para o front
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/games", controllers.GetGames)
	r.GET("games/filter", controllers.GetGamesByFilter)
	r.POST("/games", controllers.CreateGame)
	r.PUT("/games/:id", controllers.EditGame)
	r.DELETE("/games/:id", controllers.DeleteGame)

	r.Run(":8080")
}
