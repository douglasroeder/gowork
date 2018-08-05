package tests

import (
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