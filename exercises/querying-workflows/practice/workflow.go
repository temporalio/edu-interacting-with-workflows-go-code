package signals

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

type FulfillOrderSignal struct {
	Fulfilled bool
}

func Workflow(ctx workflow.Context, input string) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	logger := workflow.GetLogger(ctx)
	logger.Info("Query workflow started", "input", input)

	// TODO Part A: Add a Query Handler that returns a variable called `currentState`
	// and update the `currentState`` variable as your Workflow progresses.

	var signal FulfillOrderSignal
	var result string

	signalChan := workflow.GetSignalChannel(ctx, "fulfill-order-signal")
	signalChan.Receive(ctx, &signal)
	if signal.Fulfilled == true {
		err := workflow.ExecuteActivity(ctx, Activity, input).Get(ctx, &result)
		if err != nil {
			logger.Error("Activity failed.", "Error", err)
			return "", err
		}
		// TODO Part E: Update the `currentState` variable after the Workflow has
		// has completed, so that you can query it later.
		logger.Info("Signal workflow completed.", "result", result)
	}

	return result, nil
}

func Activity(ctx context.Context, input string) (string, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("Activity", "input", input)

	return "Received " + input, nil
}
