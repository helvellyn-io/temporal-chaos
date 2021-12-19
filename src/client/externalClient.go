package client

/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.

import (
	"context"
	"log"

	"github.com/helvellyn-io/chaos/config"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

var clientset = config.K8Client()

func GetPods() string {
	var podDeletionCandidate string
	pods, _ := clientset.CoreV1().Pods("dev").List(context.Background(),
		v1.ListOptions{LabelSelector: "app=tcx-accountservice"})

	for _, v := range pods.Items {

		podDeletionCandidate = (v.Name)
		break
	}
	return podDeletionCandidate
}

func DeletePod(p string) (string, error) {

	err := clientset.CoreV1().Pods("dev").Delete(context.TODO(), p, v1.DeleteOptions{})
	if err != nil {
		log.Fatal(err)
	}
	var result string = ("Deleted Pod: " + p)
	return result, err
}
