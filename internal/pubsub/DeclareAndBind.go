package pubsub

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"errors"
)

func DeclareAndBind(
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	queueType SimpleQueueType, // an enum to represent "durable" or "transient"
) (*amqp.Channel, amqp.Queue, error) {
	
	chan, err := conn.Channel()
	if err != nil {
		return nil, nil, errors.new("failed")
	}

	
}
