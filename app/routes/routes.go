package routes

import (
	"github.com/labstack/echo/v4"

	productController "THT-dagangan/controllers/product"
)

type ControllerList struct {
	ProductController productController.ProductController
}

func (cl *ControllerList) RouteRegister(echo *echo.Echo) {
	// product
	product := echo.Group("api/v1/product")
	product.POST("", cl.ProductController.CreateNewProduct)
	product.GET("", cl.ProductController.GetAllProducts)
	product.DELETE("/:id", cl.ProductController.DeleteProduct)
	product.PUT("/:id", cl.ProductController.UpdateProduct)
}
