package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sametxpolat/vatansoft/database"
	"github.com/sametxpolat/vatansoft/handler"
	"github.com/sametxpolat/vatansoft/repository"
	"github.com/sametxpolat/vatansoft/router"
	"github.com/sametxpolat/vatansoft/service"
)

func main() {
	e := echo.New()

	db := database.ConnectionMySQL()

	// category
	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	router.CategoryRouter(e, categoryHandler)

	// product
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	router.ProductRouter(e, productHandler)

	e.Start(":8080")
}
