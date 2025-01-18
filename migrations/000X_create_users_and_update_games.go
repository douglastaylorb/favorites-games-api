package migrations

import (
	"gorm.io/gorm"
)

func CreateUsersAndUpdateGames(db *gorm.DB) error {
	// Criar tabela de usuários
	err := db.Exec(`
        CREATE TABLE users (
            id SERIAL PRIMARY KEY,
            username VARCHAR(255) UNIQUE NOT NULL,
            email VARCHAR(255) UNIQUE NOT NULL,
            password VARCHAR(255) NOT NULL,
            created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
        )
    `).Error
	if err != nil {
		return err
	}

	// Adicionar coluna user_id à tabela de jogos
	err = db.Exec(`
        ALTER TABLE games
        ADD COLUMN user_id INTEGER REFERENCES users(id)
    `).Error
	if err != nil {
		return err
	}

	return nil
}

func UpdateGamesAddStatus(db *gorm.DB) error {
	err := db.Exec(`
			ALTER TABLE games
			ADD COLUMN status VARCHAR(50)
	`).Error
	if err != nil {
		return err
	}

	return nil
}
