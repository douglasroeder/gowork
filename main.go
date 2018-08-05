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

	categoriesController := initCategoriesController()
	router.GET("/categories", categoriesController.Index)
	router.GET("/categories/:id", categoriesController.Show)

	router.Run(":8080")
}

func initCategoriesController() controllers.CategoriesController {
	service := services.NewCategoryService(goWork.DB)

	return controllers.NewCategoriesController(service)
}

func dbAutoMigrate() {
	goWork.DB.AutoMigrate(&models.Category{})
}
