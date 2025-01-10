package models

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	Nome          string `json:"nome"`
	Genero        string `json:"genero"`
	Desenvolvedor string `json:"desenvolvedor"`
	AnoLancamento int    `json:"ano_lancamento"`
	Nota          int    `json:"nota"`
	Descricao     string `json:"descricao"`
	Imagem        string `json:"imagem"`
}
