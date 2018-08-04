package models

import "github.com/jinzhu/gorm"

// Brand stores product category
type Brand struct {
	gorm.Model
	Name string `json:"name"`
}
