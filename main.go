package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	configDB "THT-dagangan/app/configs/databases"
	_middleware "THT-dagangan/app/middlewares/logger"
	"THT-dagangan/businesses/productEntity"
	productController "THT-dagangan/controllers/product"
	_domainFactory "THT-dagangan/repository"

	"THT-dagangan/app/routes"
)

func main() {
	var (
		db = configDB.SetupDatabaseConnection()
	)
	timeoutDur, _ := strconv.Atoi(os.Getenv("CONTEXT_TIMEOUT"))
	timeoutContext := time.Duration(timeoutDur) * time.Millisecond

	echoApp := echo.New()

	//middleware
	echoApp.Use(middleware.CORS())
	echoApp.Use(middleware.LoggerWithConfig(_middleware.LoggerConfig()))

	//product
	productRepo := _domainFactory.NewProductRepository(db)
	productService := productEntity.NewProductServices(productRepo, timeoutContext)
	productCtrl := productController.NewProductController(productService)

	//routes
	routesInit := routes.ControllerList{
		ProductController: *productCtrl,
	}
	routesInit.RouteRegister(echoApp)

	port := os.Getenv("PORT")
	log.Fatal(echoApp.Start(":" + port))
}
