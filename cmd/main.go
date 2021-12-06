package main

import (
	"github.com/helvellyn-io/chaos/src/client"
	schedule "github.com/helvellyn-io/chaos/src/scheduler"
)

func main() {

	schedule.Scheduler()
	client.GetPods()

}
