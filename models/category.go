package models

import "github.com/jinzhu/gorm"

// Category stores product category
type Category struct {
	gorm.Model
	Name string `json:"name"`
}
