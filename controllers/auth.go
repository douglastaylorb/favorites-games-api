package controllers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/douglastaylorb/favorites-games-api/database"
	"github.com/douglastaylorb/favorites-games-api/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// Register cria um novo usuário
func Register(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Erro ao fazer bind do JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Dados recebidos: Username=%s, Email=%s, Senha=%s", input.Username, input.Email, input.Password)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Erro ao gerar hash da senha: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao processar senha"})
		return
	}

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		log.Printf("Erro ao criar usuário: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao criar usuário"})
		return
	}

	log.Printf("Usuário criado com sucesso: %s", user.Username)
	c.JSON(http.StatusCreated, gin.H{"message": "Usuário criado com sucesso"})
}

// Login autentica um usuário e retorna um token JWT

func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Erro ao fazer bind do JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados de login inválidos"})
		return
	}

	log.Printf("Tentativa de login para o usuário: %s", input.Username)

	var user models.User
	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		log.Printf("Usuário não encontrado: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	log.Printf("Usuário encontrado: ID=%d, Username=%s", user.ID, user.Username)
	log.Printf("Senha armazenada (hash): %s", user.Password)
	log.Printf("Senha fornecida: %s", input.Password)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		log.Printf("Erro na comparação de senhas: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	log.Printf("Autenticação bem-sucedida para o usuário: %s", user.Username)

	// Gera o token JWT
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		log.Println("AVISO: JWT_SECRET_KEY não está definida. Usando uma chave padrão insegura.")
		secretKey = "chave_padrao_insegura" // Apenas para desenvolvimento, não use em produção!
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token válido por 24 horas
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Printf("Erro ao gerar token JWT: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	log.Printf("Token JWT gerado com sucesso para o usuário: %s", user.Username)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
		"user": gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		},
	})
}
