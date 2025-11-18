package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/pubsub"
	"github.com/bootdotdev/learn-pub-sub-starter/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	connString := "amqp://guest:guest@localhost:5672/"

	con, err := amqp.Dial(connString)
	if err != nil {
		panic("Failed to setup the connection")
	}

	defer con.Close()

	c, err := con.Channel()
	if err != nil {
		panic("Failed to create amqp")
	}

	pubsub.PublishJSON(c, routing.ExchangePerilDirect, routing.PauseKey, routing.PlayingState{IsPaused: true})

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	


	fmt.Print("rabbitmq connection successfully established")

	fmt.Println("Starting Peril server...")

		for _ = range signalChan {
		fmt.Printf("Interrupt signal received. Shutting down gracefully.")
		os.Exit(0)
	}
}
