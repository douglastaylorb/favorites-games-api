package models

import (
	"time"

	"gorm.io/gorm"
)

// models/game.go
type Game struct {
	gorm.Model
	UserID        uint   `json:"user_id"`
	Nome          string `json:"nome"`
	Genero        string `json:"genero"`
	Desenvolvedor string `json:"desenvolvedor"`
	AnoLancamento int    `json:"ano_lancamento"`
	Nota          int    `json:"nota"`
	Descricao     string `json:"descricao"`
	Imagem        string `json:"imagem"`
	User          User   `json:"-" gorm:"foreignKey:UserID"`
}

// Estrutura para o Swagger
type SwaggerGame struct {
	ID            uint      `json:"id" swaggertype:"integer"`
	CreatedAt     time.Time `json:"created_at" swaggertype:"string" format:"date-time"`
	UpdatedAt     time.Time `json:"updated_at" swaggertype:"string" format:"date-time"`
	DeletedAt     time.Time `json:"deleted_at,omitempty" swaggertype:"string" format:"date-time"`
	Nome          string    `json:"nome"`
	Genero        string    `json:"genero"`
	Desenvolvedor string    `json:"desenvolvedor"`
	AnoLancamento int       `json:"ano_lancamento"`
	Nota          int       `json:"nota"`
	Descricao     string    `json:"descricao"`
	Imagem        string    `json:"imagem"`
}
