package chaos

import (
	"github.com/helvellyn-io/chaos/src/client"
)

func DeletePod() (string, error) {
	client.DeletePod(client.GetPods())

	return client.DeletePod(client.GetPods())

}
