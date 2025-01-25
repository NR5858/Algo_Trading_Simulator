package rabbitmq

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func EstablishConnection() (*amqp091.Connection, *amqp091.Channel, error) {
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	return conn, ch, nil
}

func DeclareExchange(ch *amqp091.Channel, exchangeName string) {
	err := ch.ExchangeDeclare(
		exchangeName, // name
		"direct",     // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare an exchange")
}

func DeclareQueue(ch *amqp091.Channel, queueName string) *amqp091.Queue {
	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")
	return &q
}

func BindToQueue(ch *amqp091.Channel, q *amqp091.Queue, exchangeName string) {
	err := ch.QueueBind(
		q.Name,       // queue name
		"",           // routing key
		exchangeName, // exchange
		false,
		nil)
	failOnError(err, "Failed to bind a queue")
}

func SendOrder(message string, ch *amqp091.Channel, q *amqp091.Queue) error {
	err := ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	failOnError(err, "Failed to publish a message")

	log.Printf("Sent message to RabbitMQ: %s", message)
	return nil
}
