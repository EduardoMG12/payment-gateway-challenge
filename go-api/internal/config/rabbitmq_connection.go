package config

import (
	"fmt"
	"log"
	"os"
)

type RabbitMQConnection struct {
	AmqpURI string
}

func rabbitmq_connection() *RabbitMQConnection {

	rabbitMQDefaultUser := os.Getenv("RABBITMQ_DEFAULT_USER")
	rabbitMQDefaultPass := os.Getenv("RABBITMQ_DEFAULT_PASS")
	rabbitMQHost := os.Getenv("RABBITMQ_HOST")
	rabbitMQPort := os.Getenv("RABBITMQ_PORT")

	amqpURI := "amqp://" + rabbitMQDefaultUser + ":" + rabbitMQDefaultPass + "@" + rabbitMQHost + ":" + rabbitMQPort

	if amqpURI == "" {
		log.Fatal("AMQP_URI is not configured.")
	}
	fmt.Print(amqpURI)

	return &RabbitMQConnection{
		AmqpURI: amqpURI,
	}

}
