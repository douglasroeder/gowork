package services

func (suite *TestSuite) TestGetAllOnEmptyData() {
	service := NewCategoryService(goWork.DB)

	categories := service.GetAll()
	suite.Equal(0, len(categories))
}
