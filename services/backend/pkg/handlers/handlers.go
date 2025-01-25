package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/NR5858/Algo_Trading_Simulator/pkg/models"
	"github.com/NR5858/Algo_Trading_Simulator/pkg/order"
	"github.com/NR5858/Algo_Trading_Simulator/pkg/rabbitmq"
	"github.com/rabbitmq/amqp091-go"
)

// getCurrentTimestamp returns the current timestamp with nanosecond precision
func getCurrentTimestampInMillis() string {
	return time.Now().Format("2006-01-02 15:04:05.000000")
}

func HandleOrder(w http.ResponseWriter, r *http.Request, ch *amqp091.Channel, q *amqp091.Queue) {
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
	log.Printf("[%s] Received from client: %+v", getCurrentTimestampInMillis(), o)

	fixMessage := order.ConvertToFix(&o)
	log.Printf("[%s] Converted to FIX: %s", getCurrentTimestampInMillis(), fixMessage)
	rabbitmq.SendOrder(fixMessage, ch, q)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(o)
}
