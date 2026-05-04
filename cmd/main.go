// This is the main package for the Go API server. It sets up the server, defines routes, and starts the server on port 8080.

package main

import (
	controller "go-api/Controller"

	"github.com/gin-gonic/gin"
)

//Get to test the server
func main() {
	server := gin.Default()
	server.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})


	// Create an instance of the ProductController to handle product-related requests
	ProductController := controller.NewProductController()

	// Define a route for getting products, which will be handled by the GetProducts method of the ProductController
	server.GET("/products", ProductController.GetProducts)


	//server runs on port 8080 
	server.Run(":8080")
}