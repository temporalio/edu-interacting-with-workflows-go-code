package main

import (
	"log"

	queries "interacting/exercises/querying-workflows/solution"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "queries", worker.Options{})

	w.RegisterWorkflow(queries.Workflow)
	w.RegisterActivity(queries.Activity)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
