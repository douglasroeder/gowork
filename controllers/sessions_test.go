package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/douglasroeder/gowork/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockAuthenticationService struct{}

func (a *mockAuthenticationService) AuthenticateUser(user *models.User) (*models.JwtToken, error) {
	return &models.JwtToken{Token: "abcd"}, nil
}

func TestSessions_Crete(t *testing.T) {
	service := &mockAuthenticationService{}
	controller := NewSessionsController(service)

	w := httptest.NewRecorder()
	router := gin.Default()
	router.POST("/auth", controller.Create)

	user := models.User{
		Username: "user",
		Password: "password",
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		assert.Fail(t, "Error encoding JSON")
	}

	req, _ := http.NewRequest("POST", "/auth", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var jwtToken models.JwtToken
	err = json.Unmarshal(w.Body.Bytes(), &jwtToken)
	if err != nil {
		assert.Fail(t, "Error parsing JSON response")
	}

	assert.Equal(t, "abcd", jwtToken.Token)
}
