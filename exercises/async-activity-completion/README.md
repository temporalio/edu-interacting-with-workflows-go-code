# Exercise 3: Asynchronous Activity Completion

During this exercise, you will:

- Retrieve a task token from your Activity execution
Throw an `ErrResultPending` to indicate that the Activity is waiting for an external completion.
- Use another Temporal Client to communicate the result of the asynchronous Activity back to the Workflow

Make your changes to the code in the `practice` subdirectory (look for `TODO` comments that will guide you to where you should make changes to the code). If you need a hint or want to verify your changes, look at the complete version in the `solution` subdirectory.

## Part A: Retrieving the Task Token

1. Edit the `workflow.go` file. In the `Activity()` definition, add a call to `activity.GetInfo()` that returns an `activityInfo` object.
2. From that object, extract `activityInfo.TaskToken`. To asynchronously complete an Activity, you need to store this token outside of this Workflow, so that you can call it from another Temporal Client. The most straightforward way to do this is to encode it to hexadecimal or base64 and log it to your terminal, so add a call like `logger.Info("Activity", "taskToken", hex.EncodeToString(taskToken))`.
3. Save the file.

## Part B: Set Your Activity to Return `ErrResultPending`

1. Continue editing the same Activity definition in the `workflow.go` file. You need to add a `return` statement that returns a special kind of error, `activity.ErrResultPending`, that will Temporal that the Activity has not failed but will be completed asynchronously.
2. Note that the Workflow's `StartToCloseTimeout` has been lengthened to 300 seconds for this exercise. Activities can still time out if they are running in the background.
3. Save the file.

## Part C: Configure a Client to send CompleteActivity

1. Now you can edit the `completionclient/main.go` file to call `CompleteActivity`. The first thing you'll need to do is add some way of supplying the `taskToken` specific to the Activity you are trying to complete at runtime. In a production system, you might store and retrieve the token from a database, but for now, you can configure this Client to accept it as an additional argument by adding `flag` parsing to the `main()` block:

```go
var taskToken string
flag.StringVar(&taskToken, "tasktoken", "", "Task Token of Activity to Complete")
flag.Parse()
decoded, err := hex.DecodeString(taskToken)
if err != nil {
   log.Fatalln("Unable to decode token", err)
}
```

2. Next, add the call to `CompleteActivity()`. This function requires a variable to write its result to, so provide something like `var result string`. Then, add a call to `c.CompleteActivity(context.Background(), decoded, result, err)`. Don't forget to add any error handling as needed.
3. Save the file.

## Part D: Running the Workflow and Completing it Asynchronously

At this point, you can run your Workflow. As with the Signal Exercise, the Workflow will not return on its own -- in this case, because your Activity is set to complete asynchronously, and will wait to receive `CompleteActivity()`.

1. In one terminal, navigate to the `worker` subdirectory and run `go run main.go`.
2. In another terminal, navigate to the `starter` subdirectory and run `go run main.go`. Your work will produce some logging, eventually including your `taskToken`:

```
2024/03/14 15:14:00 INFO  Activity Namespace default TaskQueue async WorkerID 22396@Omelas@ ActivityID 5 ActivityType Activity Attempt 1 WorkflowType Workflow WorkflowID async RunID 0c3cb022-042f-4437-b021-a6cf2a4afe1b taskToken 0a2461613733613533322d363337362d346130332d613563342d36626134626437306139623312056173796e631a2430633363623032322d303432662d343433372d623032312d61366366326134616665316220052801320135420841637469766974794a08080110be80401801
```

3. You can now use this token to send a `CompleteActivity()` call from another client. In a third terminal, navigate to the `completeclient` subdirectory and run `go run main.go --tasktoken [tasktoken]`, pasting the token from the previous step. This will cause your Activity to return and your Workflow to successfully complete. The terminal running your Worker process should now show `workflow completed`:

```
2024/03/14 15:15:45 INFO  Async workflow completed. Namespace default TaskQueue async WorkerID 41996@Omelas@ WorkflowType Workflow WorkflowID async RunID cf865c9f-2487-47a3-84c4-fd0c6da6ee6c Attempt 1 result
```

### This is the end of the exercise.
