package services

import (
	"fmt"

	"github.com/douglasroeder/gowork/models"
	"github.com/jinzhu/gorm"
)

// ContactService handles some of the CRUD operations
type ContactService interface {
	GetAll() []models.Contact
	GetByID(id int64) (models.Contact, bool)
	Insert(category *models.Contact) (bool, []string)
	DeleteByID(id int64) bool
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

func (s *contactService) GetByID(id int64) (models.Contact, bool) {
	var contact models.Contact

	if err := s.db.Where("id = ?", id).First(&contact).Error; err != nil {
		fmt.Println("Error")

		return contact, false
	}

	return contact, true
}

func (s *contactService) DeleteByID(id int64) bool {
	var contact models.Contact

	s.db.Where("id = ?", id).Delete(&contact)

	return true
}

func (s *contactService) Insert(contact *models.Contact) (bool, []string) {
	if s.db.NewRecord(*contact) {
		result := s.db.Create(contact)

		errors := []string{}
		for _, error := range result.GetErrors() {
			errors = append(errors, error.Error())
		}

		if len(errors) >= 1 {
			return false, errors
		}

		return true, errors
	}

	return false, []string{}
}
