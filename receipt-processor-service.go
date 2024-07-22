package main

import (
	"math"
	"strconv"
	"strings"
	"time"
)

func calculatePoints(receipt Receipt) int {
	points := 0

	// One point for every alphanumeric character in the retailer name.
	points += countAlphanumeric(receipt.Retailer)

	//50 points if the total is a round dollar amount with no cents.
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if total == float64(int(total)) {
		points += 50
	}

	//25 points if the total is a multiple of 0.25.
	if int(total*100)%25 == 0 {
		points += 25
	}

	//5 points for every two items on the receipt.
	points += 5 * (len(receipt.Items) / 2)

	//If the trimmed length of the item description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	for _, item := range receipt.Items {
		descriptionLen := len(strings.TrimSpace(item.ShortDescription))
		if descriptionLen%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	//6 points if the day in the purchase date is odd.
	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if purchaseDate.Day()%2 != 0 {
		points += 6
	}

	//10 points if the time of purchase is after 2:00pm and before 4:00pm.
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() == 14 && purchaseTime.Minute() > 0 {
		points += 10
	}

	return points
}

func countAlphanumeric(input string) int {
	count := 0
	for _, ch := range input {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') {
			count++
		}
	}
	return count
}
