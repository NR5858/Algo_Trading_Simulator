package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/NR5858/Algo_Trading_Simulator/pkg/models"
	"github.com/NR5858/Algo_Trading_Simulator/pkg/order"
)

// getCurrentTimestamp returns the current timestamp with nanosecond precision
func ogetCurrentTimestampInMillis() string {
	return time.Now().Format("2006-01-02 15:04:05.000000")
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

	var o models.Order
	err := json.NewDecoder(r.Body).Decode(&o)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("[%s] Received from client: %+v", ogetCurrentTimestampInMillis(), o)

	fixMessage := order.ConvertToFix(&o)
	log.Printf("[%s] Converted to FIX: %s", ogetCurrentTimestampInMillis(), fixMessage)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(o)
}
