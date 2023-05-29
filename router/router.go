package router

import (
	"github.com/labstack/echo/v4"
	"github.com/sametxpolat/vatansoft/handler"
)

func CategoryRouter(e *echo.Echo, categoryHandler *handler.CategoryHandler) {
	e.GET("/categories", categoryHandler.Categories)
	e.GET("/category/:id", categoryHandler.Category)
	e.POST("/category/insert", categoryHandler.Create)
	e.PUT("/category/:id/update", categoryHandler.Update)
	e.DELETE("/category/:id/delete", categoryHandler.Delete)
}

func ProductRouter(e *echo.Echo, productHandler *handler.ProductHandler) {
	e.GET("/stocks", productHandler.Stocks)
	e.GET("/stock/:id", productHandler.Stock)
	e.POST("/stock/insert", productHandler.Create)
	e.PUT("/stock/:id/update", productHandler.Update)
	e.DELETE("/stock/:id/delete", productHandler.Delete)

	// e.POST("/stocks/filter", productHandler.Filter)
}
