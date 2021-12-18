package schedule

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type ChaosSpec struct {
	TargetType string `json:"targettype"`
	Interval   int    `json:"interval"`
	RunIn      string `json:"runin"`
}

func Scheduler() interface{} {

	f, _ := ioutil.ReadFile("/Users/dylanjohnson/go/src/github.com/helvellyn-io/chaos/config/configuration.json")

	js := ChaosSpec{}
	err := json.Unmarshal(f, &js)
	if err != nil {
		os.Exit(1)
	}

	//fmt.Println(js.TargetType)

	return js.TargetType

}
