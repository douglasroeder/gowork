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

func (suite *TestSuite) TestGetByIDWhenInvalid() {
	service := NewCategoryService(goWork.DB)

	_, found := service.GetByID(9999)
	suite.Equal(false, found)
}
