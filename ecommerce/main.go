package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
        port = "8080"
    }

	app :=  controllers.NewApplication(database.ProductData(database.Client, "Products"),database.UserData(database.Client, "Users"))

    router := gin.New()
	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.Get("/addtocart",app.AddToCart())
	router.GET("/removeItem", app.RemoveItem())
	router.GET("/cartcheckout", app.BuyFromCart())
	router.GET("/instantbuy", app.Instantbuy())

	log.Fatal(router.Run(":"+port))
}