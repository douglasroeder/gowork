package services

import (
	"github.com/douglasroeder/gowork/models"
)

func (suite *TestSuite) TestCategoryService_GetAll() {
	service := NewCategoryService(goWork.DB)

	categories := service.GetAll()
	suite.Equal(2, len(categories))
}

func (suite *TestSuite) TestCategoryService_GetByIDWhenValid() {
	service := NewCategoryService(goWork.DB)

	category, found := service.GetByID(1)
	suite.Equal(true, found)
	suite.Equal("Laptops", category.Name)
}

func (suite *TestSuite) TestCategoryService_GetByIDWhenNotFound() {
	service := NewCategoryService(goWork.DB)

	_, found := service.GetByID(9999)
	suite.False(found)
}

func (suite *TestSuite) TestCategoryService_Insert() {
	service := NewCategoryService(goWork.DB)
	category := models.Category{
		Name: "Smartphone",
	}

	success, _ := service.Insert(&category)

	suite.True(success)
	suite.NotNil(category.ID)
}

func (suite *TestSuite) TestCategoryService_DeleteByID() {
	service := NewCategoryService(goWork.DB)

	deleted := service.DeleteByID(1)
	suite.True(deleted)

	_, found := service.GetByID(1)
	suite.False(found)
}
