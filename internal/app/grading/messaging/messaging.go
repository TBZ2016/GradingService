package messaging

import (
	"github.com/streadway/amqp"
)

type MessageBroker struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func NewMessageBroker(url string) (*MessageBroker, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &MessageBroker{conn, channel}, nil
}

func (mb *MessageBroker) Close() {
	mb.channel.Close()
	mb.conn.Close()
}

func (mb *MessageBroker) PublishMessage(exchange, message string) error {
	err := mb.channel.Publish(
		exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	return err
}
