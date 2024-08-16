package main

import (
	"context"
	"fmt"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := filepath.Join("/home/appscodepc/.kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	ctx := context.TODO()

	// Create a deployment
	if err := createDeployment(ctx, clientset); err != nil {
		panic(err)
	}

	// Create a service
	if err := createService(ctx, clientset); err != nil {
		panic(err)
	}

	// Create a pod
	if err := createPod(ctx, clientset); err != nil {
		panic(err)
	}

	fmt.Println("All resources created successfully.")
}
