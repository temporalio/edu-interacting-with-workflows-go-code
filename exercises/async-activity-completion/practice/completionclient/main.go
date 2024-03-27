package main

import (
	"log"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// TODO Part C: Parse the taskToken logged from your Activity as an
	// additional CLI argument to this client, then call `CompleteActivity()`
	// with that taskToken to asynchronously complete your running activity.
	// You will need to add "context", "flag", and an "encoding" import.
}
