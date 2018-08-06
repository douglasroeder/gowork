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

// IndexCategory handles /categories route
func (controller *categoriesController) Index(c *gin.Context) {
	categories := controller.service.GetAll()

	c.JSON(200, categories)
}

// ShowCategory handles /categories/1 route
func (controller *categoriesController) Show(c *gin.Context) {
	idParam := c.Params.ByName("id")
	id, _ := strconv.ParseInt(idParam, 10, 0)
	category, found := controller.service.GetByID(id)

	if found {
		c.JSON(200, category)
		return
	}

	c.JSON(404, gin.H{
		"message": "Category not found",
	})
	c.Abort()
	return
}

func (controller *categoriesController) Create(c *gin.Context) {
	var category models.Category
	if c.Bind(&category) == nil {
		success := controller.service.Insert(&category)
		if success {
			c.JSON(200, category)
			return
		}

		c.JSON(404, gin.H{
			"message": "Error creating category",
		})
		c.Abort()
		return
	}
}
