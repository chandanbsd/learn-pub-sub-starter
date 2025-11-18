package pubsub

import (
	"context"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)


func PublishJSON[T any](ch *amqp.Channel, exchange, key string, val T) error {
	valBytes, err := json.Marshal(val)
	if err != nil {
		panic("Body bytes")
	}

	pubStruct := amqp.Publishing {
		ContentType: "application/json",
		Body: valBytes,
	}

	err = ch.PublishWithContext(context.Background(), exchange, key, false, false, pubStruct)
	if err != nil {
		return err
	}
	return nil
}
