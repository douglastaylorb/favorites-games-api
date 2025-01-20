package controllers

import (
	"encoding/csv"
	"io"
	"net/http"
	"strconv"

	"github.com/douglastaylorb/favorites-games-api/database"
	"github.com/douglastaylorb/favorites-games-api/models"
	"github.com/gin-gonic/gin"
)

// Tipos de status de jogos do usuário
var validStatuses = map[string]bool{
	"Não jogado": true,
	"Jogando":    true,
	"Zerado":     true,
	"Platinado":  true,
}

var validGenres = map[string]bool{
	"Ação":          true,
	"Aventura":      true,
	"Corrida":       true,
	"Esporte":       true,
	"Estratégia":    true,
	"Luta":          true,
	"MMORPG":        true,
	"Plataforma":    true,
	"RPG":           true,
	"Simulação":     true,
	"Sobrevivência": true,
	"Terror":        true,
	"Tiro":          true,
	"Ritmo":         true,
	"Outro":         true,
}

func GetGames(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var games []models.Game
	query := database.DB

	// Parâmetros de ordenação
	sortBy := c.DefaultQuery("sort", "nome")
	order := c.DefaultQuery("order", "asc")

	// Aplica a ordenação
	switch sortBy {
	case "nome", "nota", "ano_lancamento":
		query = query.Order(sortBy + " " + order)
	default:
		query = query.Order("nome asc")
	}

	// Busca os jogos do usuário
	result := query.Where("user_id = ?", userID).Find(&games)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar jogos"})
		return
	}
	c.JSON(http.StatusOK, games)
}

// CreateGame cria um novo jogo para o usuário autenticado
func CreateGame(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var game models.Game

	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !validStatuses[game.Status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status inválido"})
		return
	}

	if !validGenres[game.Genero] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gênero inválido"})
		return
	}

	game.UserID = userID
	if err := database.DB.Create(&game).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar jogo"})
		return
	}
	c.JSON(http.StatusCreated, game)
}

// cria múltiplos jogos a partir de um arquivo CSV
func CreateGamesBulk(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Arquivo não enviado"})
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var games []models.Game

	// Pula a primeira linha (cabeçalho)
	if _, err := reader.Read(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao ler arquivo CSV"})
		return
	}

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao ler arquivo CSV"})
			return
		}

		anoLancamento, _ := strconv.Atoi(record[3])
		nota, _ := strconv.Atoi(record[4])

		game := models.Game{
			UserID:        userID,
			Nome:          record[0],
			Genero:        record[1],
			Desenvolvedor: record[2],
			AnoLancamento: anoLancamento,
			Nota:          nota,
			Status:        record[5],
			Descricao:     record[6],
			Imagem:        record[7],
		}
		games = append(games, game)
	}

	if err := database.DB.Create(&games).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar jogos no banco de dados"})
		return
	}

	c.JSON(http.StatusCreated, games)
}

// EditGame atualiza um jogo existente
func EditGame(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var game models.Game
	id := c.Param("id")

	if err := database.DB.Where("id = ? AND user_id = ?", id, userID).First(&game).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Jogo não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&game); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !validStatuses[game.Status] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status inválido"})
		return
	}

	if !validGenres[game.Genero] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gênero inválido"})
		return
	}

	if err := database.DB.Save(&game).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar jogo"})
		return
	}

	c.JSON(http.StatusOK, game)
}

// DeleteGame remove um jogo
func DeleteGame(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	id := c.Param("id")

	result := database.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Game{})
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Jogo não encontrado"})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Jogo deletado com sucesso"})
}

// GetGamesByFilter retorna jogos filtrados por ano e/ou nota mínima e/ou status
func GetGamesByFilter(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var games []models.Game
	query := database.DB.Where("user_id = ?", userID)

	if ano := c.Query("year"); ano != "" {
		query = query.Where("ano_lancamento = ?", ano)
	}

	if notaMinima := c.Query("minRating"); notaMinima != "" {
		rating, err := strconv.Atoi(notaMinima)
		if err == nil {
			query = query.Where("nota >= ?", rating)
		}
	}

	if status := c.Query("status"); status != "" {
		if _, ok := validStatuses[status]; ok {
			query = query.Where("status = ?", status)
		}
	}

	if err := query.Find(&games).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao filtrar jogos"})
		return
	}
	c.JSON(http.StatusOK, games)
}
