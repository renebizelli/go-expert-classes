package main

import (
	"renebizelli/utils/pkg/events/rabbitmq"
	"time"
)

func main() {

	ch := rabbitmq.OpenChannel()

	defer ch.Close()

	for {
		rabbitmq.Publish(ch, "amq.direct", "Hello World")
		time.Sleep(time.Second)
	}

}
