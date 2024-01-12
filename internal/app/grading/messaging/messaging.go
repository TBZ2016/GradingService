package messaging

import (
	"github.com/streadway/amqp"
)

// RabbitMQ struct to hold the connection and channel
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

// Initialize RabbitMQ connection and channel
func NewRabbitMQ(url string) (*RabbitMQ, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		conn:    conn,
		channel: channel,
	}, nil
}

func (mq *RabbitMQ) Publish(exchange, routingKey string, body []byte) error {
	err := mq.channel.Publish(
		exchange,
		routingKey,
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	return err
}

func (mq *RabbitMQ) Consume(queue, consumer string, handlerFunc func([]byte)) error {
	msgs, err := mq.channel.Consume(
		queue,
		consumer,
		true,  // auto-acknowledge
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			handlerFunc(msg.Body)
		}
	}()

	return nil
}
