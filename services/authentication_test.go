package services

import "github.com/douglasroeder/gowork/models"

func (suite *TestSuite) TestAuthenticationService_AuthenticateUser() {
	service := NewAuthenticationService()

	user := &models.User{
		Username: "username",
		Password: "password",
	}

	jwtToken, error := service.AuthenticateUser(user)
	suite.Nil(error)
	suite.NotNil(jwtToken.Token)
}
