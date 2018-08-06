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
	v1 := router.Group("v1")
	{
		categoriesGroup := v1.Group("categories")
		{
			categoriesController := initCategoriesController()
			categoriesGroup.GET("", categoriesController.Index)
			categoriesGroup.GET("/:id", categoriesController.Show)
			categoriesGroup.POST("", categoriesController.Create)
		}
	}

	router.Run(":8080")
}

func initCategoriesController() controllers.CategoriesController {
	service := services.NewCategoryService(goWork.DB)

	return controllers.NewCategoriesController(service)
}

func dbAutoMigrate() {
	goWork.DB.AutoMigrate(&models.Category{})
}
