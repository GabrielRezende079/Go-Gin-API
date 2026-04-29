package main

import "github.com/gin-gonic/gin"

func main() {

	//Get to test the server
	server := gin.Default()
	server.GET("/ping", func(ctx *gin.Context){
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//server runs on port 8080
	server.Run(":8080")
}