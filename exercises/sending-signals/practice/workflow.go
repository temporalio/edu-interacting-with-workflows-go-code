package signals

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

// TODO Part A: Define a new Signal type `struct` named `FulfillOrderSignal`.
// It should contain a single variable, a `bool`, named `Fulfilled`.

func Workflow(ctx workflow.Context, input string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("Signal workflow started", "input", input)

	var signal FulfillOrderSignal
	var result string

	signalChan := workflow.GetSignalChannel(ctx, "fulfill-order-signal")
	signalChan.Receive(ctx, &signal)
	// TODO Part B: Wrap the `ExecuteActivity()`` call and the `logger.Info()` call
	// in a test for `if signal.Fulfilled == true`. This will block the Workflow
	// until a Signal is received on the `signalChan` defined above.
	err := workflow.ExecuteActivity(ctx, Activity, input).Get(ctx, &result)
	if err != nil {
		logger.Error("Activity failed.", "Error", err)
		return "", err
	}
	logger.Info("Signal workflow completed.", "result", result)

	return result, nil
}

func Activity(ctx context.Context, input string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "input", input)

	return "Received " + input, nil
}
