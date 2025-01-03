package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	id   int64
	data string
}

func main() {

	chRabbitMQ := make(chan Message)
	chKafka := make(chan Message)

	var id int64 = 0

	//RabbitMQ
	go func() {

		for {
			atomic.AddInt64(&id, 1)
			chRabbitMQ <- Message{id, "RabbitMQ"}
			time.Sleep(1 * time.Second)
		}

	}()

	//Kafka
	go func() {

		for {
			atomic.AddInt64(&id, 1)
			chKafka <- Message{id, "Kafka"}
			time.Sleep(500 * time.Millisecond)
		}

	}()

	for {
		select {
		case msg := <-chRabbitMQ:
			fmt.Printf("Received from %s: %v\n", msg.data, msg.id)
		case msg := <-chKafka:
			fmt.Printf("Received from %s: %v\n", msg.data, msg.id)
		case <-time.After(3 * time.Second):
			fmt.Println("timeout")
		}
	}

}
