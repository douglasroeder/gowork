package controllers

import (
	"strconv"

	"github.com/douglasroeder/gowork/services"
	"github.com/gin-gonic/gin"
)

// CategoriesController manages categories
type CategoriesController interface {
	Index(c *gin.Context)
	Show(c *gin.Context)
}

// NewCategoriesController returns a new instance of CategoriesController
func NewCategoriesController(service services.CategoryService) CategoriesController {
	return &categoriesController{
		service: service,
	}
}

type categoriesController struct {
	service services.CategoryService
}

// IndexCategory handles /categories route
func (controller *categoriesController) Index(c *gin.Context) {
	categories := controller.service.GetAll()

	c.JSON(200, categories)
}

// ShowCategory handles /categories/1 route
func (controller *categoriesController) Show(c *gin.Context) {
	idParam := c.Params.ByName("id")
	id, _ := strconv.ParseInt(idParam, 10, 0)
	category, _ := controller.service.GetByID(id)

	c.JSON(200, category)
}
