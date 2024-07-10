package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	currentHour := time.Now().Hour()
	fmt.Println(currentHour)
	kubeconfig := flag.String("kubeconfig", "/Users/ayushkumar/.kube/config", "location to your kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	result, err := deploymentsClient.Get(context.TODO(), "nginx", metav1.GetOptions{})
	if err != nil {
		panic(err)
	}
	if currentHour >= 9 && currentHour < 17 {
		result.Spec.Replicas = int32Ptr(4)
	} else {
		result.Spec.Replicas = int32Ptr(2)
	}

	_, updateErr := deploymentsClient.Update(context.TODO(), result, metav1.UpdateOptions{})
	fmt.Println(updateErr)
}

func int32Ptr(i int32) *int32 { return &i }
