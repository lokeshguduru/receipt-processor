package main

import "sync"

var (
	receiptStore = make(map[string]Receipt)
	storeLock    sync.Mutex
)

// storeReceipt saves a receipt with the given ID
func storeReceipt(id string, receipt Receipt) {
	storeLock.Lock()
	defer storeLock.Unlock()
	receiptStore[id] = receipt
}

// getReceipt retrieves a receipt by ID
func getReceipt(id string) (Receipt, bool) {
	storeLock.Lock()
	defer storeLock.Unlock()
	receipt, exists := receiptStore[id]
	return receipt, exists
}