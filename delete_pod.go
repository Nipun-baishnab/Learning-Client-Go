package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := "/home/appscodepc/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %v", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err)
	}

	namespace := "default" // adjust if necessary
	podName := flag.String("pod", "", "Name of the pod to delete")
	flag.Parse()

	if *podName == "" {
		log.Fatal("Pod name must be provided")
	}

	// Deleting the specified pod
	err = clientset.CoreV1().Pods(namespace).Delete(context.TODO(), *podName, metav1.DeleteOptions{})
	if err != nil {
		log.Fatalf("Error deleting pod: %v", err)
	}

	fmt.Println("Pod deleted successfully")
}
