package controllers

import (
	"encoding/csv"
	"io"
	"net/http"
	"strconv"

	"github.com/douglastaylorb/favorites-games-api/database"
	_ "github.com/douglastaylorb/favorites-games-api/docs"
	"github.com/douglastaylorb/favorites-games-api/models"
	"github.com/gin-gonic/gin"
)

// @Summary Get all games
// @Description Get a list of all games
// @Tags games
// @Produce json
// @Param sort query string false "Sort field"
// @Param order query string false "Sort order (asc or desc)"
// @Success 200 {array} models.SwaggerGame
// @Failure 500 {object} map[string]string
// @Router /games [get]
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

func CreateGamesBulk(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Arquivo não enviado",
		})
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var games []models.Game

	_, err = reader.Read()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao ler arquivo. Verifique seu arquivo CSV.",
		})
		return
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Erro ao ler arquivo. Verifique seu arquivo CSV.",
			})
			return
		}

		anoLancamento, _ := strconv.Atoi(record[3])
		nota, _ := strconv.Atoi(record[4])

		game := models.Game{
			Nome:          record[0],
			Genero:        record[1],
			Desenvolvedor: record[2],
			AnoLancamento: anoLancamento,
			Nota:          nota,
			Descricao:     record[5],
			Imagem:        record[6],
		}
		games = append(games, game)
	}

	result := database.DB.Create(&games)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao salvar jogos no banco de dados",
		})
		return
	}

	c.JSON(http.StatusCreated, games)

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
