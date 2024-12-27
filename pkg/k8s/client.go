package k8s

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	metricsv "k8s.io/metrics/pkg/client/clientset/versioned"
)

// GetKubeClient initializes a Kubernetes clientset
func GetKubeClient(kubeconfig string) (*kubernetes.Clientset, *rest.Config, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, nil, err
	}

	return clientset, config, nil
}

func GetPodMetrics(config *rest.Config, namespace string) (map[string]map[string]string, error) {
	metricsClient, err := metricsv.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create metrics client: %w", err)
	}

	podMetrics, err := metricsClient.MetricsV1beta1().PodMetricses(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch pod metrics: %w", err)
	}

	metrics := make(map[string]map[string]string)
	for _, podMetric := range podMetrics.Items {
		metrics[podMetric.Name] = make(map[string]string)
		for _, container := range podMetric.Containers {
			metrics[podMetric.Name][container.Name] = fmt.Sprintf("CPU: %vm, Memory: %vMi", container.Usage.Cpu().MilliValue(), container.Usage.Memory().Value()/1024/1024)
		}
	}

	return metrics, nil
}
