package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

// Our struct that we will send
type Event struct {
	Name string
}

func main() {
	connection, conn_error := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/%s", "guest", "guest", "localhost:5672", ""))
	if conn_error != nil {
		panic(conn_error)
	}
	fmt.Println("Connection established", connection.RemoteAddr())

	channel, chan_error := connection.Channel()
	if chan_error != nil {
		panic(chan_error)
	}
	fmt.Println("Channel initialized:", !channel.IsClosed())

	// Start the queue
	queue, queue_error := channel.QueueDeclare("events", true, false, false, true, amqp.Table{
		// "x-queue-type":                    "stream",
		// "x-stream-max-segment-size-bytes": 30000,
		// "x-max-length-bytes":              150000
	})

	if queue_error != nil {
		panic(queue_error)
	}
	fmt.Println("Queue Created:", queue.Name)
	// Publish 1000 Messages
	ctx := context.Background()
	for i := 0; i <= 1000; i++ {

		event := Event{
			Name: "test",
		}
		data, err := json.Marshal(event)
		if err != nil {
			panic(err)
		}

		err = channel.PublishWithContext(ctx, "", "events", false, false, amqp.Publishing{
			CorrelationId: uuid.NewString(),
			Body:          data,
		})
		if err != nil {
			panic(err)
		}
	}
	channel.Close()
}
