package tests

import "github.com/douglasroeder/gowork/services"

func (suite *TestSuite) TestCategoryService_GetAll() {
	service := services.NewCategoryService(goWork.DB)

	categories := service.GetAll()
	suite.Equal(2, len(categories))
}

func (suite *TestSuite) TestCategoryService_GetByIDWhenValid() {
	service := services.NewCategoryService(goWork.DB)

	category, found := service.GetByID(1)
	suite.Equal(true, found)
	suite.Equal("Laptops", category.Name)
}

func (suite *TestSuite) TestCategoryService_GetByIDWhenNotFound() {
	service := services.NewCategoryService(goWork.DB)

	_, found := service.GetByID(9999)
	suite.False(found)
}

func (suite *TestSuite) TestCategoryService_DeleteByID() {
	service := services.NewCategoryService(goWork.DB)

	deleted := service.DeleteByID(1)
	suite.True(deleted)

	_, found := service.GetByID(1)
	suite.False(found)
}
