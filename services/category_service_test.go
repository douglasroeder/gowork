package services

func (suite *TestSuite) TestGetAll() {
	service := NewCategoryService(goWork.DB)

	categories := service.GetAll()
	suite.Equal(2, len(categories))
}

func (suite *TestSuite) TestGetByIDWhenValid() {
	service := NewCategoryService(goWork.DB)

	category, found := service.GetByID(1)
	suite.Equal(true, found)
	suite.Equal("Laptops", category.Name)
}

func (suite *TestSuite) TestGetByIDWhenNotFound() {
	service := NewCategoryService(goWork.DB)

	_, found := service.GetByID(9999)
	suite.False(found)
}

func (suite *TestSuite) TestDeleteByID() {
	service := NewCategoryService(goWork.DB)

	deleted := service.DeleteByID(1)
	suite.True(deleted)

	_, found := service.GetByID(1)
	suite.False(found)
}
