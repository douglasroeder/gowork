package services

import "github.com/douglasroeder/gowork/models"

func (suite *TestSuite) TestAuthenticationService_AuthenticateUser() {
	service := NewAuthenticationService()

	validUser := &models.User{
		Username: "username",
		Password: "1234",
	}

	jwtToken, error := service.AuthenticateUser(validUser)
	suite.Nil(error)
	suite.NotNil(jwtToken.Token)

	invalidUser := &models.User{
		Username: "username",
		Password: "wrong_password",
	}
	jwtToken, error = service.AuthenticateUser(invalidUser)
	suite.NotNil(error)
	suite.Nil(jwtToken)
}
