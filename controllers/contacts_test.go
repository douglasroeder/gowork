package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/douglasroeder/gowork/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockContactService struct {
	categories []models.Contact
}

func (s *mockContactService) GetAll() []models.Contact {
	contacts := []models.Contact{
		models.Contact{ID: 1, Name: "John Doe"},
		models.Contact{ID: 2, Name: "Jane Doe"},
	}
	return contacts
}

func TestContacts_Index(t *testing.T) {
	service := &mockContactService{}
	controller := NewContactsController(service)

	w := httptest.NewRecorder()
	router := gin.Default()
	router.GET("/index", controller.Index)

	req, _ := http.NewRequest("GET", "/index", nil)
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var result models.Result
	err := json.Unmarshal(w.Body.Bytes(), &result)
	if err != nil {
		assert.Fail(t, "Error parsing JSON response")
	}

	payload := result.Payload.([]interface{})
	assert.Equal(t, 2, len(payload))
	assert.Equal(t, "John Doe", payload[0].(map[string]interface{})["name"])
}
