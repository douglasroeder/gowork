package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/douglasroeder/gowork/controllers"
	"github.com/douglasroeder/gowork/models"
	"github.com/douglasroeder/gowork/services"

	"github.com/gin-gonic/gin"
)

func (suite *TestSuite) TestCategoriesController_Index() {
	service := services.NewCategoryService(goWork.DB)
	controller := controllers.NewCategoriesController(service)

	w := httptest.NewRecorder()
	router := gin.Default()
	router.GET("/index", controller.Index)

	req, _ := http.NewRequest("GET", "/index", nil)
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	suite.Equal(200, w.Code)

	var categories []models.Category
	err := json.Unmarshal(w.Body.Bytes(), &categories)
	if err != nil {
		suite.Fail("Error parsing JSON response")
	}

	suite.Equal(2, len(categories))
	suite.Equal("Laptops", categories[0].Name)
}

func (suite *TestSuite) TestCategoriesController_ShowWhenFound() {
	service := services.NewCategoryService(goWork.DB)
	controller := controllers.NewCategoriesController(service)

	w := httptest.NewRecorder()
	router := gin.Default()
	router.GET("/show/:id", controller.Show)

	req, _ := http.NewRequest("GET", "/show/1", nil)
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	suite.Equal(200, w.Code)

	var category models.Category
	err := json.Unmarshal(w.Body.Bytes(), &category)
	if err != nil {
		suite.Fail("Error parsing JSON response")
	}

	suite.Equal("Laptops", category.Name)
}

func (suite *TestSuite) TestCategoriesController_ShowWhenNotFound() {
	service := services.NewCategoryService(goWork.DB)
	controller := controllers.NewCategoriesController(service)

	w := httptest.NewRecorder()
	router := gin.Default()
	router.GET("/show/:id", controller.Show)

	req, _ := http.NewRequest("GET", "/show/999", nil)
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	suite.Equal(404, w.Code)

	var response = struct {
		Message string `json:"message"`
	}{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		suite.Fail("Error parsing JSON response")
	}

	suite.Equal("Category not found", response.Message)
}

func (suite *TestSuite) TestCategoriesController_Create() {
	service := services.NewCategoryService(goWork.DB)
	controller := controllers.NewCategoriesController(service)

	w := httptest.NewRecorder()
	router := gin.Default()
	router.POST("/create", controller.Create)

	category := models.Category{
		Name: "Smartphone",
	}

	jsonData, err := json.Marshal(category)
	if err != nil {
		suite.Fail("Error encoding JSON")
	}

	req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(jsonData))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	var createdCategory models.Category
	err = json.Unmarshal(w.Body.Bytes(), &createdCategory)
	if err != nil {
		suite.Fail("Error parsing JSON response")
	}

	suite.Equal(200, w.Code)
	suite.NotNil(createdCategory.ID)
	suite.Equal("Smartphone", createdCategory.Name)
}
