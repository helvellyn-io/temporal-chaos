package main

import (
	"fmt"
	"os"

	"github.com/helvellyn-io/chaos/src/client"
)

func main() {
	r, err := client.DeletePod(client.GetPods())
	if r == "" {
		os.Exit(1)

	}
	fmt.Println(r, err)

}
