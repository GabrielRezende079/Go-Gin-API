/*
product_controller.go is responsible for handling product-related requests and responses in the API.
It defines a ProductController struct and a method GetProducts to retrieve a list of products and return them as JSON.
*/

package controller

import (
	"go-api/model"

	"github.com/gin-gonic/gin"
)

// ProductController handles product-related requests and responses. It defines methods to retrieve product data and return it in JSON format.
type ProductController struct {
}

// NewProductController creates and returns a new instance of ProductController.
func NewProductController() ProductController {
	return ProductController{}
}

/* 
GetProducts retrieves a list of products and returns them as mocked JSON. 
It creates a slice of Product structs with sample data and sends it in the response with a 200 status code.
*/
func (p *ProductController) GetProducts(ctx *gin.Context) {
	products := []model.Product{
		{ID: 1, Name: "Fries", Price: 10.0},
		{ID: 2, Name: "Burger", Price: 15.0},
	}
	ctx.JSON(200, products)
	}
