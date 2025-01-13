package main

import (
	"log"
	// TODO Part C: Add signals "interacting/exercises/sending-signals-client/practice"
	// to your module imports.
	// TODO Part D: Add "context" to your module imports.

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	// TODO Part C: Use the `FullfillOrderSignal` struct type from the `signals` module
	// (i.e., the `workflow.go` file in the parent directory).
	// Create an instance of `FulfillOrderSignal` that contains `Fulfilled: true`.

	// TODO Part D: Call `SignalWorkflow()` to send a Signal to your running Workflow.
	// It needs, as arguments, `context.Background()`, your workflow ID, your run ID
	// (which can be an empty string), the name of the signal, and the signal instance.
	// It should assign its result to `err` so that it can be checked in the next line.
	if err != nil {
		log.Fatalln("Error sending the Signal", err)
		return
	}
}
