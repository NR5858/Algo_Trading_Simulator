package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Order struct {
	Symbol    string `json:"symbol"`
	Action    string `json:"action"`
	Quantity  int    `json:"quantity"`
	PriceType string `json:"priceType"`
	Duration  string `json:"duration"`
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	// Add CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var order Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Process the order (e.g., save to database, send to message queue, etc.)
	// For now, just log the order
	log.Printf("Received from client: %+v", order)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}
