package database

import (
	"fmt"
	"log"
	"os"

	"github.com/douglastaylorb/favorites-games-api/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	// Carrega as variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}

	// Obtém as credenciais do banco de dados das variáveis de ambiente
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Monta a string de conexão
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Conecta ao banco de dados
	DB, err = gorm.Open(postgres.Open(connection))
	if err != nil {
		log.Panic("Erro ao conectar ao banco de dados: ", err)
	}

	// Executa as migrações automáticas
	DB.AutoMigrate(&models.User{}, &models.Game{})

	log.Println("Conexão com o banco de dados estabelecida com sucesso")
}
