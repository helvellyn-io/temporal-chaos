package main

import (
	"log"

	"github.com/helvellyn-io/chaos"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.NewClient(client.Options{
		HostPort: "20.122.110.191:7233", //client.DefaultHostPort,
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "cron", worker.Options{})

	w.RegisterWorkflow(chaos.DelPodWorkflow)
	w.RegisterActivity(chaos.DelPod)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
