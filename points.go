package main

import (
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// calculatePoints computes the points for a receipt based on the rules
func calculatePoints(receipt Receipt) int {
	points := 0

	// Rule 1: One point per alphanumeric character in retailer name
	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			points++
		}
	}

	// Rule 2: 50 points if total is a round dollar amount
	if strings.HasSuffix(receipt.Total, ".00") {
		points += 50
	}

	// Rule 3: 25 points if total is a multiple of 0.25
	if cents, err := strconv.Atoi(receipt.Total[len(receipt.Total)-2:]); err == nil && cents%25 == 0 {
		points += 25
	}

	// Rule 4: 5 points for every two items
	points += (len(receipt.Items) / 2) * 5

	// Rule 5: Points for items with description length multiple of 3
	for _, item := range receipt.Items {
		trimmedDesc := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDesc)%3 == 0 && len(trimmedDesc) > 0 {
			if price, err := strconv.ParseFloat(item.Price, 64); err == nil {
				points += int(math.Ceil(price * 0.2))
			}
		}
	}

	// Rule 6: 6 points if purchase day is odd
	if date, err := time.Parse("2006-01-02", receipt.PurchaseDate); err == nil && date.Day()%2 == 1 {
		points += 6
	}

	// Rule 7: 10 points if purchase time is between 2:00pm and 4:00pm
	if t, err := time.Parse("15:04", receipt.PurchaseTime); err == nil {
		hour, minute := t.Hour(), t.Minute()
		if (hour == 14 && minute > 0) || (hour == 15) {
			points += 10
		}
	}

	return points
}