package main

import (
	"fmt"
	"os"
	"os/signal"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	connString := "amqp://guest:guest@localhost:5672/"

	con, err := amqp.Dial(connString)
	if err != nil {
		panic("Failed to setup the connection")
	}

	defer con.Close()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	
	for _ = range signalChan {
		fmt.Printf("Interrupt signal received. Shutting down gracefully.")
		os.Exit(0)
	}

	fmt.Print("rabbitmq connection successfully established")

	fmt.Println("Starting Peril server...")
}
