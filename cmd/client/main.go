package main

import (
	"fmt"

	"github.com/bootdotdev/learn-pub-sub-starter/internal/gamelogic"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril client...")
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

	username, err := gamelogic.ClientWelcome()
	if err != nil {
		fmt.Println("")
		return
	}

}
