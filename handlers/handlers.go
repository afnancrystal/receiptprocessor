package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"receiptprocessor/models"
	"receiptprocessor/store"
	"receiptprocessor/utils"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Calculate and log points
	points := utils.CalculatePoints(receipt)
	// ID generation
	id := uuid.New().String()

	// Save points in memory
	store.Mu.Lock()
	store.Receipts[id] = points
	store.Mu.Unlock()

	// Log to terminal
	fmt.Printf("🧾 Stored receipt - ID: %s | Points: %d\n", id, points)

	// Return ID as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"id": id,
	})
}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	store.Mu.Lock()
	points, exists := store.Receipts[id]
	store.Mu.Unlock()

	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	// Return points as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{
		"points": points,
	})
}
