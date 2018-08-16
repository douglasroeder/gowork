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

func (s *mockContactService) GetByID(id int64) (models.Contact, bool) {
	if id == 1 {
		contact := models.Contact{ID: 1, Name: "John Doe"}
		return contact, true
	}

	return models.Contact{}, false
}

func (s *mockContactService) Insert(contact *models.Contact) (bool, []string) {
	return true, []string{}
}

func (s *mockContactService) DeleteByID(id int64) bool {
	return true
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

func TestContacts_ShowWhenFound(t *testing.T) {
	service := &mockContactService{}
	controller := NewContactsController(service)

	w := httptest.NewRecorder()
	router := gin.Default()
	router.GET("/show/:id", controller.Show)

	req, _ := http.NewRequest("GET", "/show/1", nil)
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	var result models.Result
	err := json.Unmarshal(w.Body.Bytes(), &result)
	if err != nil {
		assert.Fail(t, "Error parsing JSON response")
	}

	assert.Equal(t, "John Doe", result.Payload.(map[string]interface{})["name"])
}

func TestContacts_ShowWhenNotFound(t *testing.T) {
	service := &mockContactService{}
	controller := NewContactsController(service)

	w := httptest.NewRecorder()
	router := gin.Default()
	router.GET("/show/:id", controller.Show)

	req, _ := http.NewRequest("GET", "/show/999", nil)
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)

	var result models.Result
	err := json.Unmarshal(w.Body.Bytes(), &result)
	if err != nil {
		assert.Fail(t, "Error parsing JSON response")
	}

	assert.Equal(t, "Category not found", result.Error)
}

func TestContacts_Create(t *testing.T) {
	service := &mockContactService{}
	controller := NewContactsController(service)

	w := httptest.NewRecorder()
	router := gin.Default()
	router.POST("/create", controller.Create)

	contact := models.Contact{
		Name: "John Doe",
	}

	jsonData, err := json.Marshal(contact)
	if err != nil {
		assert.Fail(t, "Error encoding JSON")
	}

	req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Add more meaninful test coverage here
	assert.Equal(t, 200, w.Code)
}
