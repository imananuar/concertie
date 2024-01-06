package rmq

import (
	"context"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/joho/godotenv"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func PublishMessage(queueName string, data string) {

	if err := godotenv.Load(); err != nil {
		log.Panicf("Error loading .env file: %s", err)
		return
	}

	HOSTNAME := os.Getenv("RMQ_HOSTNAME")
	USERNAME := os.Getenv("RMQ_USERNAME")
	PASSWORD := os.Getenv("RMQ_PASSWORD")

	RMQ_URL := "amqp://" + USERNAME + ":" + PASSWORD + "@" + HOSTNAME

	conn, err := amqp.Dial(RMQ_URL)
	failOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "failed to open a channel")
	// defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	  )
	  failOnError(err, "Failed to declare a queue")
	  
	  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	  defer cancel()
	  
	//   body := "{'pattern': 'purchase_order', 'data': 'purchased!'}"
		body := "{\"pattern\": \"purchase_order\", \"data\": \"" + data + "\"}"

	  err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
		  ContentType: "text/plain",
		  Body:        []byte(body),
		})
	  failOnError(err, "Failed to publish a message")
	  log.Printf(" [x] Sent %s\n", body)
}