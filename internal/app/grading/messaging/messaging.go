// messaging/rabbitmq.go
package messaging

import (
	"fmt"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn *amqp.Connection
}

func NewRabbitMQ(connectionString string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	return &RabbitMQ{conn: conn}, nil
}

func (rmq *RabbitMQ) Close() {
	if rmq.conn != nil {
		rmq.conn.Close()
	}
}

// Setup declares an exchange, a queue, and binds them together.
func (rmq *RabbitMQ) Setup(exchangeName, queueName, routingKey string) error {
	ch, err := rmq.conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare an exchange
	err = ch.ExchangeDeclare(
		exchangeName,
		"direct", // or "fanout" or "topic" depending on your needs
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declare an exchange: %v", err)
	}

	
	_, err = ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to declare a queue: %v", err)
	}

	err = ch.QueueBind(
		queueName,
		routingKey,
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to bind queue to exchange: %v", err)
	}

	return nil
}

// SendMessage publishes a message to the RabbitMQ exchange.
func (rmq *RabbitMQ) SendMessage(message []byte, exchangeName, routingKey string) error {
	ch, err := rmq.conn.Channel()
	if err != nil {
		return fmt.Errorf("failed to open a channel: %v", err)
	}
	defer ch.Close()

	err = ch.Publish(
		exchangeName,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish a message: %v", err)
	}

	return nil
}

// ConsumeMessages consumes messages from the RabbitMQ queue.
func (rmq *RabbitMQ) ConsumeMessages(queueName, consumerName string) (<-chan amqp.Delivery, error) {
	ch, err := rmq.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %v", err)
	}

	messages, err := ch.Consume(
		queueName,
		consumerName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to consume messages: %v", err)
	}

	return messages, nil
}
