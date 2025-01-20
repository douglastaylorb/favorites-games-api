package routes

import (
	"github.com/douglastaylorb/favorites-games-api/controllers"
	middlewares "github.com/douglastaylorb/favorites-games-api/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	// Config cors para o front
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 60 * 60, // 12 horas
	}))

	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", controllers.Login)
	}

	api := r.Group("/api")
	api.Use(middlewares.AuthMiddleware())
	{
		api.GET("/games", controllers.GetGames)
		api.GET("/games/filter", controllers.GetGamesByFilter)
		api.POST("/games", controllers.CreateGame)
		api.POST("/games/bulk", controllers.CreateGamesBulk)
		api.PUT("/games/:id", controllers.EditGame)
		api.DELETE("/games/:id", controllers.DeleteGame)
	}

	r.Run(":8080")
}
