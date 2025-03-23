package consumer

import (
	"Work2Rabbit/internal/dbconn"
	"Work2Rabbit/internal/rabbitconn"
	"context"
	"log"
)

func ProcessQueueRed() {
	processQueue(func(s string) {
		log.Printf("Red message: %s", s)
		c := context.Background()
		_, err := dbconn.DB.CreateRedWord(c, s)
		if err != nil {
			log.Printf("Send error: %v", err)
			panic(err)
		}
	})
}

func ProcessQueueGreen() {
	processQueue(func(s string) {
		log.Printf("Green message: %s", s)
		c := context.Background()
		dbconn.DB.CreateGreenWord(c, s)
	})
}

func processQueue(processMessage func(string)) {
	q, err := rabbitconn.Channel.QueueDeclare(
		"words", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Printf("Declare err: %v", err)
		panic(err)
	}
	err = rabbitconn.Channel.ExchangeDeclare(
		"wordsEx", // name
		"direct",  // type
		true,      // durable
		false,     // auto-deleted
		false,     // internal
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Printf("Exchange err: %v", err)
		panic(err)
	}
	msgs, _ := rabbitconn.Channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	func() {
		for d := range msgs {
			processMessage(string(d.Body))
		}
	}()
}
