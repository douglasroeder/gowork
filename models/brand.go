package models

import "github.com/jinzhu/gorm"

// Brand stores product category
type Brand struct {
	gorm.Model
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
