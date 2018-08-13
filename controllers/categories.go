package controllers

import (
	"strconv"

	"github.com/douglasroeder/gowork/models"
	"github.com/douglasroeder/gowork/services"
	"github.com/gin-gonic/gin"
)

// CategoriesController manages categories
type CategoriesController interface {
	Index(c *gin.Context)
	Show(c *gin.Context)
	Create(c *gin.Context)
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

// Index handles GET /categories route
func (controller *categoriesController) Index(c *gin.Context) {
	categories := controller.service.GetAll()
	response := models.NewResult(200, categories, []string{})
	c.JSON(200, response)
}

// Show handles GET /categories/1 route
func (controller *categoriesController) Show(c *gin.Context) {
	idParam := c.Params.ByName("id")
	id, _ := strconv.ParseInt(idParam, 10, 0)
	category, found := controller.service.GetByID(id)

	if found {
		c.JSON(200, models.NewResult(200, category, []string{}))
		return
	}

	c.JSON(404, models.NewResult(404, nil, []string{"Category not found"}))
	c.Abort()
	return
}

// Create handles POST /categories route
func (controller *categoriesController) Create(c *gin.Context) {
	var category models.Category
	if c.BindJSON(&category) == nil {
		success, errors := controller.service.Insert(&category)
		if success {
			c.JSON(200, category)
			return
		}

		c.JSON(404, gin.H{
			"message": "Error creating category",
			"errors":  errors,
		})
		c.Abort()
		return
	}
}
