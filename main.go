package main

import (
	"github.com/douglasroeder/gowork/app"
	"github.com/douglasroeder/gowork/controllers"
	"github.com/douglasroeder/gowork/models"
	"github.com/douglasroeder/gowork/services"
	"github.com/gin-gonic/gin"
)

var goWork *app.App

func main() {
	goWork = app.NewApp()
	defer goWork.Close()

	// migrating tables
	dbAutoMigrate()

	router := gin.Default()
	router.Use(CORSMiddleware())

	sessionsController := initSessionsController()
	router.POST("/authenticate", sessionsController.Create)

	v1 := router.Group("v1")
	{
		categoriesGroup := v1.Group("categories")
		{
			categoriesController := initCategoriesController()
			categoriesGroup.GET("", categoriesController.Index)
			categoriesGroup.GET("/:id", categoriesController.Show)
			categoriesGroup.POST("", categoriesController.Create)
		}

		contactsGroup := v1.Group("contacts")
		{
			contactsController := initContactsController()
			contactsGroup.GET("", contactsController.Index)
			contactsGroup.GET("/:id", contactsController.Show)
			contactsGroup.POST("", contactsController.Create)
		}
	}

	router.Run(":8080")
}

func initSessionsController() controllers.SessionsController {
	service := services.NewAuthenticationService()

	return controllers.NewSessionsController(service)
}

func initCategoriesController() controllers.CategoriesController {
	service := services.NewCategoryService(goWork.DB)

	return controllers.NewCategoriesController(service)
}

func initContactsController() controllers.ContactsController {
	service := services.NewContactService(goWork.DB)

	return controllers.NewContactsController(service)
}

func dbAutoMigrate() {
	goWork.DB.AutoMigrate(&models.Category{}, &models.Contact{})
}

// CORSMiddleware allows cors
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
