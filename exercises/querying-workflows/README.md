# Exercise #2: Querying Workflows

During this exercise, you will:

- Define and handle a Query
- Create a handle on a Workflow to be queried
- Call the Query from the Client
- Send a Query from the Command line

Make your changes to the code in the `practice` subdirectory (look for
`TODO` comments that will guide you to where you should make changes to
the code). If you need a hint or want to verify your changes, look at
the complete version in the `solution` subdirectory.

## Part A: Defining a Query

In this part of the exercise, you will define your Query.

1. Edit the `workflow.go` file. You will use the `workflow.SetQueryHandler()` to add Query handling to your Workflow. This can be done anywhere inside of the `Workflow()` function. In this case, you'll add a query to return the Workflow's `currentState`, so it makes sense to initialize the `currentState` variable as you go along. Note that this is the same Workflow from the previous exercise, so it is designed to wait for a Signal after starting.
2. Set up two string variables, `currentState := "started"` and `queryType := "current_state"`. One of these will contain the type of your Query, which will not change. The other will contain the `currentState` that is progressively updated and will be returned by the Query.
3. Next, add the `workflow.SetQueryHandler()` function with the necessary error handling:

```go
err := workflow.SetQueryHandler(ctx, queryType, func() (string, error) {
   return currentState, nil
})
if err != nil {
   currentState = "failed to register query handler"
   return "", err
}
```

4. Finally, update `currentState` again after that block of code, to something like "waiting for signal". This is the state that the Workflow will be waiting at when we query it.
5. Save the file.

## Part B: Performing a Query from a Client

In this part of the exercise, you will create another Temporal client that sends a Query. `queryclient/main.go` currently contains a near-empty Temporal client -- it looks like a Starter or a Worker, but it does not register any Workflows or do anything. You will use this client to send a Query.

1. Edit the `queryclient/main.go` file. You will use `client.QueryWorkflow(context, "workflow-id", "run-id", "query-type")` to send a Query.
2. Within the `main()` block, add the `client.QueryWorkflow()` function with the necessary parameters and error handling:

```go
response, err := c.QueryWorkflow(context.Background(), "queries", "", "current_state")
if err != nil {
   log.Fatalln("Error sending the Query", err)
   return
}
```

3. Immediately after that, add some logging for getting the Query result and logging it to your terminal:

```go
var result string
response.Get(&result)
log.Println("Received Query result. Result: " + result)
```

4. Save the file.

## Part C: Running the Workflow and the Query

At this point, you can run your Workflow. Because it is the same Workflow from the last exercise with added Query support, it will wait to receive a Signal after starting.

1. In one terminal, navigate to the `worker` subdirectory and run `go run main.go`.
2. In another terminal, navigate to the `starter` subdirectory and run `go run main.go`. You should receive some logging from your Worker along these lines:

```
2024/03/14 08:48:10 INFO  No logger configured for temporal client. Created default one.
2024/03/14 08:48:10 INFO  Started Worker Namespace default TaskQueue queries WorkerID 43388@Omelas@
2024/03/14 08:48:21 INFO  Query workflow started Namespace default TaskQueue queries WorkerID 43388@Omelas@ WorkflowType Workflow WorkflowID queries RunID 905330da-9c0f-490e-bd48-c6a9e8840f7a Attempt 1 input Plain text input
```

3. You can now Query your Workflow. In a third terminal, navigate to the `queryclient` subdirectory and run `go run main.go`. It will send a Query to your Workflow, which will immediately return the Query result:

```
2024/03/14 10:32:09 Received Query result. Result: waiting for signal
```

You have now demonstrated sending a Query to a Workflow from a Temporal Client. Before moving on, you can send one from the CLI as well. 

## Part D: Sending a Query from the Command Line

To send a Query from the CLI, use `temporal workflow query` with the same parameters as your client:

```bash
temporal workflow query \
    --workflow-id="queries" \
    --type="current_state"
```

It will produce the same result:

```
Query result:
["waiting for signal"]
```

Now you can send a Signal to your Workflow as in the previous exercise so it completes successfully.

## Part E: Querying a Completed Workflow

Finally, you can demonstrate querying completed Workflows. Update the "currentState" variable in `workflow.go` once more just before the Workflow returns, so that you can demonstrate querying a completed workflow. Then, re-run the workflow, and query it from the command line again:

```bash
temporal workflow query \
    --workflow-id="queries" \
    --type="current_state"
```

```
Query result:
["workflow completed"]
```

You need to have a running Worker to query completed Workflows.

### This is the end of the exercise.
