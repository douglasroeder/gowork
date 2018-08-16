package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/douglasroeder/gowork/models"
	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

type mockCategoryService struct {
	categories []models.Category
}

func (s *mockCategoryService) GetAll() []models.Category {
	categories := []models.Category{
		models.Category{ID: 1, Name: "Laptops"},
		models.Category{ID: 2, Name: "Memory"},
	}
	return categories
}

func (s *mockCategoryService) GetByID(id int64) (models.Category, bool) {
	if id == 1 {
		category := models.Category{ID: 1, Name: "Laptops"}
		return category, true
	}

	return models.Category{}, false
}

func (s *mockCategoryService) Insert(category *models.Category) (bool, []string) {
	return true, []string{}
}

func (s *mockCategoryService) DeleteByID(id int64) bool {
	return true
}

func TestCategories_Index(t *testing.T) {
	service := &mockCategoryService{}
	controller := NewCategoriesController(service)

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
	assert.Equal(t, "Laptops", payload[0].(map[string]interface{})["name"])
}

func TestCategories_ShowWhenFound(t *testing.T) {
	service := &mockCategoryService{}
	controller := NewCategoriesController(service)

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

	assert.Equal(t, "Laptops", result.Payload.(map[string]interface{})["name"])
}

func TestCategories_ShowWhenNotFound(t *testing.T) {
	service := &mockCategoryService{}
	controller := NewCategoriesController(service)

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

func TestCategories_Create(t *testing.T) {
	service := &mockCategoryService{}
	controller := NewCategoriesController(service)

	w := httptest.NewRecorder()
	router := gin.Default()
	router.POST("/create", controller.Create)

	category := models.Category{
		Name: "Smartphone",
	}

	jsonData, err := json.Marshal(category)
	if err != nil {
		assert.Fail(t, "Error encoding JSON")
	}

	req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	// Add more meaninful test coverage here
	assert.Equal(t, 200, w.Code)
}
