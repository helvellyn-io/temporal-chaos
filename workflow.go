package chaos

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/helvellyn-io/chaos/src/client"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

// CronResult is used to return data from one cron run to the next
type CronResult struct {
	RunTime time.Time
}

// SampleCronWorkflow executes on the given schedule
// The schedule is provided when starting the Workflow
func DelPodWorkflow(ctx workflow.Context) (*CronResult, error) {

	workflow.GetLogger(ctx).Info("Cron workflow started.", "StartTime", workflow.Now(ctx))

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx1 := workflow.WithActivityOptions(ctx, ao)

	// Start from 0 for first cron job
	lastRunTime := time.Time{}
	// Check to see if there was a previous cron job
	if workflow.HasLastCompletionResult(ctx) {
		var lastResult CronResult
		if err := workflow.GetLastCompletionResult(ctx, &lastResult); err == nil {
			lastRunTime = lastResult.RunTime
		}
	}
	thisRunTime := workflow.Now(ctx)

	err := workflow.ExecuteActivity(ctx1, DelPod, lastRunTime, thisRunTime).Get(ctx, nil)
	if err != nil {
		// Cron job failed
		// Next cron will still be scheduled by the Server
		workflow.GetLogger(ctx).Error("Cron job failed.", "Error", err)
		return nil, err
	}

	return &CronResult{RunTime: thisRunTime}, nil
}

// DoSomething is an Activity
func DelPod(ctx context.Context, lastRunTime, thisRunTime time.Time) error {
	activity.GetLogger(ctx).Info("Cron job running.", "lastRunTime_exclude", lastRunTime, "thisRunTime_include", thisRunTime)

	r := client.DeletePod(client.GetPods())
	if r == "" {
		os.Exit(1)

	}
	fmt.Println(r)
	return nil
}
