package services

import "github.com/douglasroeder/gowork/models"

func (suite *TestSuite) TestContactService_GetAll() {
	service := NewContactService(goWork.DB)

	contacts := service.GetAll()
	suite.Equal(2, len(contacts))
}

func (suite *TestSuite) TestContactService_GetByIDWhenValid() {
	service := NewContactService(goWork.DB)

	contact, found := service.GetByID(1)
	suite.Equal(true, found)
	suite.Equal("John Doe", contact.Name)
}

func (suite *TestSuite) TestContactService_GetByIDWhenNotFound() {
	service := NewContactService(goWork.DB)

	_, found := service.GetByID(9999)
	suite.False(found)
}

func (suite *TestSuite) TestContactService_Insert() {
	service := NewContactService(goWork.DB)
	contact := models.Contact{
		Name: "Mary Ann",
	}

	success, _ := service.Insert(&contact)

	suite.True(success)
	suite.NotNil(contact.ID)
}

func (suite *TestSuite) TestContactService_DeleteByID() {
	service := NewContactService(goWork.DB)

	deleted := service.DeleteByID(1)
	suite.True(deleted)

	_, found := service.GetByID(1)
	suite.False(found)
}
