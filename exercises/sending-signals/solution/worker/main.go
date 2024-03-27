package main

import (
	"log"

	signals "interacting/exercises/sending-signals/solution"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "signals", worker.Options{})

	w.RegisterWorkflow(signals.Workflow)
	w.RegisterActivity(signals.Activity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
