package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

// processReceiptHandler handles POST /receipts/process
func processReceiptHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost || r.URL.Path != "/receipts/process" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	var receipt Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := validateReceipt(receipt); err != nil {
		http.Error(w, "The receipt is invalid.", http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	storeReceipt(id, receipt)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

// getPointsHandler handles GET /receipts/{id}/points
func getPointsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet || !strings.HasPrefix(r.URL.Path, "/receipts/") {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	pathParts := strings.Split(strings.TrimPrefix(r.URL.Path, "/receipts/"), "/")
	if len(pathParts) != 2 || pathParts[1] != "points" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	id := pathParts[0]
	receipt, found := getReceipt(id)
	if !found {
		http.Error(w, "No receipt found for that ID.", http.StatusNotFound)
		return
	}

	points := calculatePoints(receipt)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"points": points})
}

func main() {
	http.HandleFunc("/receipts/process", processReceiptHandler)
	http.HandleFunc("/receipts/", getPointsHandler)
	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}