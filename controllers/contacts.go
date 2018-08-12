package controllers

import (
	"github.com/douglasroeder/gowork/models"
	"github.com/douglasroeder/gowork/services"
	"github.com/gin-gonic/gin"
)

// ContactsController manages contacts
type ContactsController interface {
	Index(c *gin.Context)
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
	response := models.NewResult(200, contacts, []string{})
	c.JSON(200, response)
}
