// filepath: /path/to/monorepo/backend/cmd/app/main.go
package main

import (
	"log"
	"net/http"

	"github.com/NR5858/Algo_Trading_Simulator/pkg/handlers"
	"github.com/NR5858/Algo_Trading_Simulator/pkg/rabbitmq"
)

func main() {
	conn, ch, err := rabbitmq.EstablishConnection()
	if err != nil {
		log.Fatalf("Error establishing RabbitMQ connection: %v", err)
	}
	defer conn.Close()
	defer ch.Close()

	const exchangeName string = "fix_exchange"
	const queueName string = "fix_messages"

	rabbitmq.DeclareExchange(ch, exchangeName)
	q := rabbitmq.DeclareQueue(ch, queueName)
	rabbitmq.BindToQueue(ch, q, exchangeName)

	http.HandleFunc("/api/order", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandleOrder(w, r, ch, q)
	})

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
