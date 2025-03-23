package rabbitconn

import (
	"Work2Rabbit/internal/config"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var Channel *amqp.Channel

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Init(conf *config.Config) {
	rabInfo := fmt.Sprintf("amqp://%s:%s@%s:%d/",
		conf.RabbitUser, conf.RabbitPassword, conf.RabbitHost, conf.RabbitPort)

	conn, err := amqp.Dial(rabInfo)
	failOnError(err, "Failed to connect to RabbitMQ")

	Channel, err = conn.Channel()
	failOnError(err, "Failed to open a Channelannel")
}
