package models

import "github.com/jinzhu/gorm"

// Contact stores contact details
type Contact struct {
	gorm.Model
	Name        string `json:"name"`
	Email       string `json:"email"`
	LandPhone   string `json:"land_phone"`
	MobilePhone string `json:"mobile_phone"`
	CnpjCpf     string `json:"cnpj_cpf"`
	IeRg        string `json:"ie_rg"`
	ContactType string `json:"contact_type"`
	Address     string `json:"address"`
	Suburb      string `json:"suburb"`
	PostalCode  string `json:"postal_code"`
	City        string `json:"city"`
	State       string `json:"state"`
	CountryCode string `json:"country_code"`
}
