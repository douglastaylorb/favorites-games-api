// models/user.go
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"size:60" json:"-"` // Especifica o tamanho do campo para 60 caracteres // O "-" impede que a senha seja serializada para JSON
	Games    []Game `json:"games"`
}
