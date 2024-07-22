package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var receiptStore = make(map[string]Receipt)

func ProcessReceipt(context *gin.Context) {
	var receipt Receipt
	if err := context.BindJSON(&receipt); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	receipt.ID = uuid.New().String()
	receiptStore[receipt.ID] = receipt

	response := map[string]string{"id": receipt.ID}
	context.JSON(http.StatusOK, response)
}

func GetPoints(context *gin.Context) {
	id := context.Param("id")

	receipt, exists := receiptStore[id]
	if !exists {
		context.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
		return
	}

	points := calculatePoints(receipt)
	response := map[string]int{"points": points}
	context.JSON(http.StatusOK, response)
}
