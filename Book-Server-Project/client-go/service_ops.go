package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func createService(ctx context.Context, clientset *kubernetes.Clientset) error {
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: "book-server-service",
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app": "book-server",
			},
			Ports: []corev1.ServicePort{
				{
					Protocol:   corev1.ProtocolTCP,
					Port:       80,
					TargetPort: int32Ptr(8080),
				},
			},
			Type: corev1.ServiceTypeLoadBalancer,
		},
	}

	servicesClient := clientset.CoreV1().Services(corev1.NamespaceDefault)
	fmt.Println("Creating service...")
	resultSvc, err := servicesClient.Create(ctx, service, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	fmt.Printf("Created service %q.\n", resultSvc.GetObjectMeta().GetName())
	return nil
}
