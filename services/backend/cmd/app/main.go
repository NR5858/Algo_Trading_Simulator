// filepath: /path/to/monorepo/backend/cmd/app/main.go
package main

import (
	"log"
	"net/http"

	"github.com/NR5858/Algo_Trading_Simulator/pkg/handlers"
)

func main() {
	http.HandleFunc("/api/order", handlers.CreateOrder)
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
