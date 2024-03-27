package main

import (
	"context"
	"log"

	queries "interacting/exercises/querying-workflows/solution"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	signal := queries.FulfillOrderSignal{
		Fulfilled: true,
	}

	err = c.SignalWorkflow(context.Background(), "queries", "", "fulfill-order-signal", signal)
	if err != nil {
		log.Fatalln("Error sending the Signal", err)
		return
	}
}
