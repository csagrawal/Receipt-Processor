package main

import (
	"receipt-processor/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func testServer(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", testServer)
	router.Run()
}
