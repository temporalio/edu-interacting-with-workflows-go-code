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

	// TODO Part B: Add the QueryWorkflow() call and log the result.
	// Don't forget to add "context" to your imports.
}
