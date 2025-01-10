package controllers

import (
	"net/http"
	"strconv"

	"github.com/douglastaylorb/favorites-games-api/database"
	"github.com/douglastaylorb/favorites-games-api/models"

	"github.com/gin-gonic/gin"
)

func GetGames(c *gin.Context) {
	var games []models.Game
	query := database.DB

	// ordenation parameter
	sortBy := c.DefaultQuery("sort", "nome")
	order := c.DefaultQuery("order", "asc")

	switch sortBy {
	case "nome":
		query = query.Order("nome " + order)
	case "nota":
		query = query.Order("nota " + order)
	case "ano_lancamento":
		query = query.Order("ano_lancamento " + order)
	default:
		query = query.Order("nome asc")
	}

	result := query.Find(&games)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao buscar jogos",
		})
		return
	}
	c.JSON(200, games)
}

func CreateGame(c *gin.Context) {
	var game models.Game

	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&game)
	c.JSON(http.StatusCreated, game)
}

func EditGame(c *gin.Context) {
	var game models.Game
	id := c.Param("id")
	database.DB.First(&game, id)

	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Model(&game).UpdateColumns(&game)
	c.JSON(http.StatusOK, game)
}

func DeleteGame(c *gin.Context) {
	var game models.Game
	id := c.Param("id")
	database.DB.First(&game, id)
	database.DB.Delete(&game)
	c.JSON(http.StatusNoContent, gin.H{
		"message": "Game deletado com sucesso",
	})
}

func GetGamesByFilter(c *gin.Context) {
	var games []models.Game
	query := database.DB

	if ano := c.Query("year"); ano != "" {
		query = query.Where("ano_lancamento = ?", ano)
	}

	if notaMinima := c.Query("minRating"); notaMinima != "" {
		rating, err := strconv.Atoi(notaMinima)
		if err == nil {
			query = query.Where("nota >= ?", rating)
		}
	}

	result := query.Find(&games)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao filtrar jogos",
		})
		return
	}
	c.JSON(http.StatusOK, games)
}
