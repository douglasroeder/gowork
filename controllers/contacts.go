package controllers

import (
	"strconv"

	"github.com/douglasroeder/gowork/models"
	"github.com/douglasroeder/gowork/services"
	"github.com/gin-gonic/gin"
)

// ContactsController manages contacts
type ContactsController interface {
	Index(c *gin.Context)
	Show(c *gin.Context)
	Create(c *gin.Context)
}

// NewContactsController returns a new instance of ContactsController
func NewContactsController(service services.ContactService) ContactsController {
	return &contactsController{
		service: service,
	}
}

type contactsController struct {
	service services.ContactService
}

// IndexCategory handles /categories route
func (controller *contactsController) Index(c *gin.Context) {
	contacts := controller.service.GetAll()
	response := models.NewResult(200, contacts, "")
	c.JSON(200, response)
}

// Show handles GET /contact/1 route
func (controller *contactsController) Show(c *gin.Context) {
	idParam := c.Params.ByName("id")
	id, _ := strconv.ParseInt(idParam, 10, 0)
	contact, found := controller.service.GetByID(id)

	if found {
		c.JSON(200, models.NewResult(200, contact, ""))
		return
	}

	c.JSON(404, models.NewResult(404, nil, "Category not found"))
	c.Abort()
	return
}

// Create handles POST /contact route
func (controller *contactsController) Create(c *gin.Context) {
	var contact models.Contact
	if c.BindJSON(&contact) == nil {
		success, _ := controller.service.Insert(&contact)
		if success {
			c.JSON(200, models.NewResult(200, contact, ""))
			return
		}

		c.JSON(404, models.NewResult(200, nil, "Error creating category"))
		c.Abort()
		return
	}
}
