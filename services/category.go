package services

import (
	"fmt"

	"github.com/douglasroeder/gowork/models"
	"github.com/jinzhu/gorm"
)

// CategoryService handles some of the CRUD operations
type CategoryService interface {
	GetAll() []models.Category
	GetByID(id int64) (models.Category, bool)
	Insert(category *models.Category) (bool, []string)
	DeleteByID(id int64) bool
}

// NewCategoryService returns and instance of CategoryService
func NewCategoryService(db *gorm.DB) CategoryService {
	return &categoryService{
		db: db,
	}
}

type categoryService struct {
	db *gorm.DB
}

func (s *categoryService) GetAll() []models.Category {
	var categories []models.Category

	if err := s.db.Find(&categories).Error; err != nil {
		fmt.Println("Error")
	}

	return categories
}

func (s *categoryService) GetByID(id int64) (models.Category, bool) {
	var category models.Category

	if err := s.db.Where("id = ?", id).First(&category).Error; err != nil {
		fmt.Println("Error")

		return category, false
	}

	return category, true
}

func (s *categoryService) DeleteByID(id int64) bool {
	var category models.Category

	s.db.Where("id = ?", id).Delete(&category)

	return true
}

func (s *categoryService) Insert(category *models.Category) (bool, []string) {
	if s.db.NewRecord(*category) {
		result := s.db.Create(category)

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
