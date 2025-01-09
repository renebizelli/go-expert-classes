package main

import (
	"fmt"
	"renebizelli/utils/pkg/events/rabbitmq"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	ch := rabbitmq.OpenChannel()

	defer ch.Close()

	msgs := make(chan amqp.Delivery)

	go rabbitmq.Consumer(ch, msgs, "orders")

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}

}
