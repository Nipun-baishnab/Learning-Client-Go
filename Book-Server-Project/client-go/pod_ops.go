package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func createPod(ctx context.Context, clientset *kubernetes.Clientset) error {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "book-server-pod",
			Labels: map[string]string{
				"app": "book-server",
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "book-server",
					Image: "arnabbaishnab7620/book-server:01",
					Ports: []corev1.ContainerPort{
						{
							ContainerPort: 8080,
						},
					},
				},
			},
		},
	}

	podsClient := clientset.CoreV1().Pods(corev1.NamespaceDefault)
	fmt.Println("Creating pod...")
	result, err := podsClient.Create(ctx, pod, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("Created pod %q.\n", result.GetObjectMeta().GetName())
	return nil
}

func deletePod(ctx context.Context, clientset *kubernetes.Clientset, podName string) error {
	podsClient := clientset.CoreV1().Pods(corev1.NamespaceDefault)
	fmt.Println("Deleting pod...")
	err := podsClient.Delete(ctx, podName, metav1.DeleteOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("Deleted pod %q.\n", podName)
	return nil
}
