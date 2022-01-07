package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/helvellyn-io/chaos/config"
	"github.com/helvellyn-io/chaos/src/redis"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var clientset = config.K8Client()

type ChangeLog struct {
	Epoch      time.Time `json:"timestamp"`
	Deployment string    `json:"deployment"`
	Namespace  string    `json:"namespace"`
	DeletedPod string    `json:"deletedPod"`
}

func GetPods() string {
	var podDeletionCandidate string
	pods, _ := clientset.CoreV1().Pods("kube-system").List(context.Background(), //namespace
		v1.ListOptions{LabelSelector: "k8s-app=kube-dns"}) //labels to filter on

	for _, v := range pods.Items {
		podDeletionCandidate = (v.Name)
		break
	}
	return podDeletionCandidate
}

func DeletePod(p string) (string, error) {
	err := clientset.CoreV1().Pods("kube-system").Delete(context.TODO(), p, v1.DeleteOptions{})
	if err != nil {
		log.Fatal(err)
	}

	l := ChangeLog{
		Deployment: "k8s-app=kube-dns",
		Namespace:  "kube-system",
		DeletedPod: p,
	}

	json, err := json.Marshal(l)
	if err != nil {
		fmt.Println(err)
	}

	redis.RedisClient().Set(redis.GenUUID(), json, 0).Err()
	var result string = ("Deleted Pod: " + p)
	return result, err
}
