package temporalweather

import (
	"github.com/helvellyn-io/chaos/src/client"
)

func DeletePod() (string, error) {
	var e error
	client.DeletePod(client.GetPods())

	return client.DeletePod(client.GetPods()), e

}