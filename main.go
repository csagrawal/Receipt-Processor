package main

import (
	"receipt-processor/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	router := gin.Default()
	router.POST("/receipts/process", ProcessReceipt)
	router.GET("/receipts/:id/points", GetPoints)

	router.Run()
}
