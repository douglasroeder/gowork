package services

import (
	"fmt"

	"github.com/douglasroeder/gowork/models"
	"github.com/jinzhu/gorm"
)

// ContactService handles some of the CRUD operations
type ContactService interface {
	GetAll() []models.Contact
}

// NewContactService returns and instance of ContactService
func NewContactService(db *gorm.DB) ContactService {
	return &contactService{
		db: db,
	}
}

type contactService struct {
	db *gorm.DB
}

func (s *contactService) GetAll() []models.Contact {
	var contacts []models.Contact

	if err := s.db.Find(&contacts).Error; err != nil {
		fmt.Println("Error")
	}

	return contacts
}
