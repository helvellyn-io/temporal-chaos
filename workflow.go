package temporalweather

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

func DeletePodWorkflow(ctx workflow.Context) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)
	var result string
	err := workflow.ExecuteActivity(ctx, DeletePod).Get(ctx, &result)
	return result, err
}
