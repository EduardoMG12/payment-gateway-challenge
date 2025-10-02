package connection

import (
	"context"
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQClient interface {
	Publish(ctx context.Context, queueName string, message []byte) error
}

type rabbitMQClientImpl struct {
	conn *amqp091.Connection
}

func NewRabbitMQClient(amqpURI string) (RabbitMQClient, error) {
	conn, err := amqp091.Dial(amqpURI)
	if err != nil {
		return nil, fmt.Errorf("fail to connect RabbitMQ: %w", err)
	}

	return &rabbitMQClientImpl{conn: conn}, nil
}

func (c *rabbitMQClientImpl) Publish(ctx context.Context, queueName string, message []byte) error {
	ch, err := c.conn.Channel()
	if err != nil {
		return fmt.Errorf("fail to open channel: %w", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("fail to declare queue: %w", err)
	}

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	if err != nil {
		return fmt.Errorf("fail to publish message: %w", err)
	}

	return nil
}
