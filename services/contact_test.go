package services

func (suite *TestSuite) TestContactService_GetAll() {
	service := NewContactService(goWork.DB)

	contacts := service.GetAll()
	suite.Equal(2, len(contacts))
}
