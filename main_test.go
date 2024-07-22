package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/receipts/process", ProcessReceipt)
	router.GET("/receipts/:id/points", GetPoints)
	return router
}

func TestProcessReceipt(test *testing.T) {
	router := setupRouter()

	receipt := Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "Klarbrunn 12-PK 12 FL OZ", Price: "12.00"},
		},
		Total: "35.35",
	}

	jsonValue, _ := json.Marshal(receipt)
	req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, req)

	assert.Equal(test, http.StatusOK, writer.Code)

	var response map[string]string
	err := json.Unmarshal(writer.Body.Bytes(), &response)
	assert.NoError(test, err)
	assert.NotEmpty(test, response["id"])
}

func TestGetPoints(test *testing.T) {
	router := setupRouter()

	receipt := Receipt{
		Retailer:     "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Items: []Item{
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
		},
		Total: "9.00",
	}

	jsonValue, _ := json.Marshal(receipt)
	req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	writer := httptest.NewRecorder()
	router.ServeHTTP(writer, req)

	var response map[string]string
	err := json.Unmarshal(writer.Body.Bytes(), &response)
	assert.NoError(test, err)
	assert.NotEmpty(test, response["id"])

	id := response["id"]

	req, _ = http.NewRequest("GET", "/receipts/"+id+"/points", nil)
	writer = httptest.NewRecorder()
	router.ServeHTTP(writer, req)

	assert.Equal(test, http.StatusOK, writer.Code)

	var pointsResponse map[string]int
	err = json.Unmarshal(writer.Body.Bytes(), &pointsResponse)
	assert.NoError(test, err)
	assert.Equal(test, 109, pointsResponse["points"])
}
