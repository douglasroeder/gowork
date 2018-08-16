package controllers

import (
	"github.com/douglasroeder/gowork/models"
	"github.com/douglasroeder/gowork/services"
	"github.com/gin-gonic/gin"
)

// SessionsController manages user sessions
type SessionsController interface {
	Create(c *gin.Context)
}

// NewSessionsController returns a new instance of SessionsController
func NewSessionsController(service services.AuthenticationService) SessionsController {
	return &sessionsController{
		service: service,
	}
}

type sessionsController struct {
	service services.AuthenticationService
}

func (controller *sessionsController) Create(c *gin.Context) {
	var user models.User
	if c.BindJSON(&user) == nil {
		jwtToken, error := controller.service.AuthenticateUser(&user)

		if error != nil {
			c.JSON(404, models.NewResult(404, nil, error.Error()))
			c.Abort()
			return
		}

		c.JSON(200, models.NewResult(200, jwtToken, ""))
		return
	}
}
