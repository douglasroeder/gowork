package models

import "time"

// Contact stores contact details
type Contact struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	LandPhone   string    `json:"land_phone"`
	MobilePhone string    `json:"mobile_phone"`
	CnpjCpf     string    `json:"cnpj_cpf"`
	IeRg        string    `json:"ie_rg"`
	ContactType string    `json:"contact_type"`
	Address     string    `json:"address"`
	Suburb      string    `json:"suburb"`
	PostalCode  string    `json:"postal_code"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	CountryCode string    `json:"country_code"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
