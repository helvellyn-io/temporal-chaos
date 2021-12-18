package main

import (
	"context"
	"fmt"
	"log"

	"go.temporal.io/sdk/client"

	app "github.com/helvellyn-io/chaos"
)

func main() {
	// Client object just once per process
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()
	options := client.StartWorkflowOptions{
		ID:        "chaos-pod-workflow",
		TaskQueue: app.ChaosMonkey,
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, app.DeletePodWorkflow)
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}
	fmt.Println(we)
}
