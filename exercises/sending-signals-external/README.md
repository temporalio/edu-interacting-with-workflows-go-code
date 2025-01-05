## Exercise #1: Sending an External Signal

During this exercise, you will:

- Define and handle a Signal
- Retrieve a handle on the Workflow to Signal
- Send an external Signal
- Use a Temporal Client to submit execution requests for both Workflows

Make your changes to the code in the `practice` subdirectory (look for
`TODO` comments that will guide you to where you should make changes to
the code). If you need a hint or want to verify your changes, look at
the complete version in the `solution` subdirectory.

## Part A: Defining a Signal

This exercise contains one Client that runs two different Workflows
— `PizzaWorkflow` and `FulfillOrderWorkflow`. Both Workflows are defined in
`workflow.go`. `PizzaWorkflow` is designed not to complete its final activity
— `SendBill` — until it receives a Signal from `FulfillOrderWorkflow`. You'll
start by defining that Signal. 

1. Edit `workflow.go`. Near the top of the file,
   after the `import()` block and before your Workflow definitions, create a
   type of `struct{}` named `FulfillOrderSignal` that contains a single `bool`
   named `Fulfilled`.
2. Save the file.

## Part B: Handling the Signal

Next, you need to enable your `PizzaWorkflow` to receive a Signal from
`FulfillOrderWorkflow`. 

After `var confirmation OrderConfirmation`, define a
Signal Channel, and use `signalChan.Receive()` to block the Workflow until it
receives a Signal, after which it can proceed with the logic contained in
`if receivedSignal.Fulfilled == true{}`. 

1. Create a `var` named `receivedSignal` that is an instance
   of `FulfillOrderSignal`. This is the Signal that
   `PizzaWorkflow` will receive from `FulfillOrderWorkflow`.
2. Add `signalChan := workflow.GetSignalChannel(ctx, "fulfill-order-signal")`.
3. On the following line, add `signalChan.Receive(ctx, &receivedSignal)`.
4. Save the file.

## Part C: Signaling your Workflow

Near the bottom of `workflow.go`, within `FulfillOrderWorkflow`, you will
notice that it runs two Activities — `MakePizzas` and `DeliverPizzas`. After
those Activities complete successfuly, the next step should be to send a
Signal to the `PizzaWorkflow` that it is time to bill the customer and
complete the Workflow. To do this, you need call
`workflow.SignalExternalWorkflow()`.

1. Create a `var` named `signalToSend` that is an instance
   of `FulfillOrderSignal` with `Fulfilled: true`. This is the Signal that
   `FulfillOrderWorkflow` will send to `PizzaWorkflow`.
2. Add a call to `workflow.SignalExternalWorkflow()` at the end of `FulfillOrderWorkflow`. `SignalExternalWorkflow`
   needs the following as arguments: the `ctx` Workflow context, Workflow ID (which should be
   `pizza-workflow-order-Z1238`), an optional Run ID (which you can omit by
   providing "" as the argument), and the name of the
   Signal,`fulfill-order-signal`. For `SignalExternalWorkflow` calls to block
   and return properly in Go, you also need to append `.Get(ctx,
   [return-value-pointer])` to a `SignalExternalWorkflow` call, though
   `[return-value-pointer]` can be `nil` here. 
   @@@@@ the course content doesn't say anything about this `.Get`. @@@@
3. Save and close the file.

## Part D: Making your Client start both Workflows

1. Finally, open `start/main.go` for editing.

Currently, this Client only starts
the `PizzaWorkflow`.

2. Directly after the `c.ExecuteWorkflow()` call for the
   `PizzaWorkflow`, add another call that starts the `FulfillOrderWorkflow`. You
   can use the call that starts the `PizzaWorkflow` and the
   `fulfillOrderWorkflowOptions` block as a reference. Don't forget to capture the
   Workflow Execution and any errors in different variables.
2. Save and close the file.

## Part E: Running both Workflows

At this point, you can run your Workflows.

1. In one terminal, navigate to the `worker` subdirectory and run `go run main.go`.
2. In another terminal, navigate to the `start` subdirectory and run `go run
   main.go`. You should receive output from both Workflows having started and
   returning the expected result:

   ```
   2024/04/12 10:41:11 Started workflow WorkflowID pizza-workflow-order-Z1238 RunID d41177c4-ffe6-4f51-a884-7fefb3e13cff
   2024/04/12 10:41:11 Started workflow WorkflowID signal-fulfilled-order-Z1238 RunID 9d5168e9-1c58-41c3-aca2-ea8182dff11d
   2024/04/12 10:41:11 Workflow result: {
     "OrderNumber": "Z1238",
     "Status": "SUCCESS",
     "ConfirmationNumber": "AB9923",
     "BillingTimestamp": 1712943671,
     "Amount": 2700
   }
   ```

3. If you look at the terminal running your Worker, you should see logging from
   each individual step run by both Workflows, including the Signal being sent
   and all the related activities:

   ```
   ...
   2024/04/12 10:41:11 INFO  Starting delivery Namespace default TaskQueue pizza-tasks WorkerID 35880@ted.local@ ActivityID 11 ActivityType DeliverPizzas Attempt 1 WorkflowType FulfillOrderWorkflow WorkflowID signal-fulfilled-order-Z1238 RunID 9d5168e9-1c58-41c3-aca2-ea8182dff11d Z1238 to {701 Mission Street Apartment 9C San Francisco CA 94103}
   2024/04/12 10:41:11 INFO  Z1238 Namespace default TaskQueue pizza-tasks WorkerID 35880@ted.local@ ActivityID 11 ActivityType DeliverPizzas Attempt 1 WorkflowType FulfillOrderWorkflow WorkflowID signal-fulfilled-order-Z1238 RunID 9d5168e9-1c58-41c3-aca2-ea8182dff11d delivered.
   ...
   ```

### This is the end of the exercise.
