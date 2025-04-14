package main

import (
	"errors"
	"regexp"
	"time"
)

// validateReceipt ensures the receipt conforms to the API spec
func validateReceipt(receipt Receipt) error {
	// Check retailer
	if receipt.Retailer == "" || !regexp.MustCompile(`^[\w\s\-&]+$`).MatchString(receipt.Retailer) {
		return errors.New("invalid retailer: must be non-empty and match allowed characters")
	}

	// Check purchaseDate
	if receipt.PurchaseDate == "" {
		return errors.New("purchaseDate is required")
	}
	if _, err := time.Parse("2006-01-02", receipt.PurchaseDate); err != nil {
		return errors.New("invalid purchaseDate: must be in YYYY-MM-DD format")
	}

	// Check purchaseTime
	if receipt.PurchaseTime == "" {
		return errors.New("purchaseTime is required")
	}
	if _, err := time.Parse("15:04", receipt.PurchaseTime); err != nil {
		return errors.New("invalid purchaseTime: must be in HH:MM (24-hour) format")
	}

	// Check items
	if len(receipt.Items) < 1 {
		return errors.New("items must contain at least one entry")
	}
	for i, item := range receipt.Items {
		if item.ShortDescription == "" || !regexp.MustCompile(`^[\w\s\-]+$`).MatchString(item.ShortDescription) {
			return errors.New("invalid shortDescription in item " + string(rune(i)) + ": must be non-empty and match allowed characters")
		}
		if !regexp.MustCompile(`^\d+\.\d{2}$`).MatchString(item.Price) {
			return errors.New("invalid price in item " + string(rune(i)) + ": must be in format N.NN")
		}
	}

	// Check total
	if !regexp.MustCompile(`^\d+\.\d{2}$`).MatchString(receipt.Total) {
		return errors.New("invalid total: must be in format N.NN")
	}

	return nil
}