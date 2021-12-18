# chaos
Random Chaos Stuff 


Define and schedule chaos experients. 


Spec: 

```func GetSvcPods() {

	svc, _ := clientset.CoreV1().Services("dev").List(context.Background(),
		v1.ListOptions{LabelSelector: "app=tcx-accountservice"})

	for _, v := range svc.Items {
		fmt.Println(v.Name)
	}

}```