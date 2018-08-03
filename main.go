package main

import (
	"fmt"

	"github.com/douglasroeder/gowork/controllers"
	"github.com/douglasroeder/gowork/models"
	"github.com/douglasroeder/gowork/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

func main() {
	db, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// migrating tables
	dbAutoMigrate()

	router := gin.Default()

	categoriesController := initCategoriesController()
	router.GET("/categories", categoriesController.IndexCategory)
	router.GET("/categories/:id", categoriesController.ShowCategory)

	router.Run(":8080")
}

func initCategoriesController() controllers.CategoriesController {
	service := services.NewCategoryService(db)

	return controllers.NewCategoriesController(service)
}

func dbAutoMigrate() {
	db.AutoMigrate(&models.Category{})
}
