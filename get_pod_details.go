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
	// Define command-line flags
	podName := flag.String("pod", "", "Name of the pod")
	namespace := flag.String("namespace", "default", "Namespace of the pod")
	kubeconfig := flag.String("kubeconfig", "/home/appscodepc/.kube/config", "Path to kubeconfig file")
	flag.Parse()

	// Validate the podName is not empty
	if *podName == "" {
		log.Fatalf("The -pod flag is required and cannot be empty")
	}

	// Load kubeconfig and create client
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %v", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating kubernetes client: %v", err)
	}

	// Get pod details
	pod, err := clientset.CoreV1().Pods(*namespace).Get(context.TODO(), *podName, metav1.GetOptions{})
	if err != nil {
		log.Fatalf("Error getting pod: %v", err)
	}

	// Print pod details
	fmt.Printf("Pod Name: %s\n", pod.Name)
	fmt.Printf("Pod Namespace: %s\n", pod.Namespace)
	fmt.Printf("Pod Status: %s\n", pod.Status.Phase)
}
