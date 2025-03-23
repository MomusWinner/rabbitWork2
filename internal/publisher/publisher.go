package publisher

import (
	"Work2Rabbit/internal/config"
	"Work2Rabbit/internal/rabbitconn"
	"bufio"
	"context"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
	"time"
)

func ProcessInput(conf *config.Config) error {
	q, err := rabbitconn.Channel.QueueDeclare(
		"words", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Print(err)
		return err
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
		log.Print(err)
		return err
	}

	err = rabbitconn.Channel.QueueBind(
		q.Name,    // queue name
		q.Name,    // routing key
		"wordsEx", // exchange
		false,
		nil,
	)

	if err != nil {
		log.Print(err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filePath := conf.InputFile

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		err = rabbitconn.Channel.PublishWithContext(ctx,
			"wordsEx", // exChannelange
			q.Name,    // routing key
			false,     // mandatory
			false,     // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(line),
			})
		if err != nil {
			log.Print(err)
			return err
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return err
}
