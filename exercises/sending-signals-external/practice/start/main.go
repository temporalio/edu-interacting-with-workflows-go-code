package main

import (
	"context"
	"encoding/json"
	"fmt"
	pizza "interacting/exercises/sending-signals-external/practice"
	"log"

	"go.temporal.io/sdk/client"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	order := *createPizzaOrder()

	pizzaWorkflowID := fmt.Sprintf("pizza-workflow-order-%s", order.OrderNumber)
	signalFulfilledID := fmt.Sprintf("signal-fulfilled-order-%s", order.OrderNumber)

	pizzaWorkflowOptions := client.StartWorkflowOptions{
		ID:        pizzaWorkflowID,
		TaskQueue: pizza.TaskQueueName,
	}

	signalFulfilledOptions := client.StartWorkflowOptions{
		ID:        signalFulfilledID,
		TaskQueue: pizza.TaskQueueName,
	}

	we, err := c.ExecuteWorkflow(context.Background(), pizzaWorkflowOptions, pizza.PizzaWorkflow, order)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}
	log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())

	// TODO Part D: Add another call to c.ExecuteWorkflow() that spawns the
	// `FulfillOrderWorkflow`. You can use the call that spawns the `PizzaWorkflow`
	// and the `signalFulfilledOptions` block above as a reference.
	// Don't forget to capture the Workflow Execution and any errors in different variables.

	var result pizza.OrderConfirmation
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable to get workflow result", err)
	}

	data, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		log.Fatalln("Unable to format order confirmation as JSON", err)
	}
	log.Printf("Workflow result: %s\n", string(data))
}

func createPizzaOrder() *pizza.PizzaOrder {
	customer := pizza.Customer{
		CustomerID: 12983,
		Name:       "María García",
		Email:      "maria1985@example.com",
		Phone:      "415-555-7418",
	}

	address := pizza.Address{
		Line1:      "701 Mission Street",
		Line2:      "Apartment 9C",
		City:       "San Francisco",
		State:      "CA",
		PostalCode: "94103",
	}

	p1 := pizza.Pizza{
		Description: "Large, with mushrooms and onions",
		Price:       1500,
	}

	p2 := pizza.Pizza{
		Description: "Small, with pepperoni",
		Price:       1200,
	}

	items := []pizza.Pizza{p1, p2}

	order := pizza.PizzaOrder{
		OrderNumber: "Z1238",
		Customer:    customer,
		Items:       items,
		Address:     address,
		IsDelivery:  true,
	}

	return &order
}
