package main

import (
	"Work2Rabbit/internal/config"
	"Work2Rabbit/internal/consumer"
	"Work2Rabbit/internal/dbconn"
	"Work2Rabbit/internal/publisher"
	"Work2Rabbit/internal/rabbitconn"
)

func main() {
	conf := config.LoadConfig(".")
	dbconn.Init(&conf)
	rabbitconn.Init(&conf)

	var forever chan struct{}

	go publisher.ProcessInput(&conf)

	go consumer.ProcessQueueRed()
	go consumer.ProcessQueueGreen()

	<-forever
}
