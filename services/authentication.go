package services

import (
	"errors"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/douglasroeder/gowork/models"
)

// AuthenticationService handles user authentication
type AuthenticationService interface {
	AuthenticateUser(user *models.User) (*models.JwtToken, error)
}

// NewAuthenticationService returns an AuthenticationService instance
func NewAuthenticationService() AuthenticationService {
	return &authenticationService{}
}

type authenticationService struct{}

func (s *authenticationService) AuthenticateUser(user *models.User) (*models.JwtToken, error) {
	// TODO validate user somewhere
	if user.Password != "1234" {
		return nil, errors.New("error authenticating user")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
	})

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		return nil, error
	}

	return &models.JwtToken{Token: tokenString}, nil
}
