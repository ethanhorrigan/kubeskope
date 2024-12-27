package cmd

import (
	"fmt"
	"github.com/ethanhorrigan/kubeskope/pkg/analysis"
	"github.com/ethanhorrigan/kubeskope/pkg/k8s"
	"github.com/ethanhorrigan/kubeskope/pkg/output"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

var namespace string
var kubeconfig string

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor resource usage in a namespace",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Monitoring namespace: %s\n", namespace)

		client, config, err := k8s.GetKubeClient(kubeconfig)
		if err != nil {
			fmt.Printf("Error getting Kubernetes client: %v\n", err)
			return
		}

		pods, err := client.CoreV1().Pods(namespace).List(cmd.Context(), metav1.ListOptions{})
		if err != nil {
			log.Fatalf("Failed to list pods: %v\n", err)
			return
		}

		fmt.Printf("Pods in namespace %s:\n", namespace)

		usageMetrics, err := k8s.GetPodMetrics(config, namespace)
		if err != nil {
			log.Fatalf("Error fetching usage metrics: %v", err)
		}

		var rows [][]string
		threshold := 90.0

		for _, pod := range pods.Items {
			fmt.Printf("- %s\n", pod.Name)
			podRows := analysis.AnalyzePod(pod, usageMetrics, threshold)
			rows = append(rows, podRows...)
		}

		output.RenderTable(rows)
	},
}

func init() {
	rootCmd.AddCommand(monitorCmd)
	monitorCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "Namespace to monitor")
	monitorCmd.Flags().StringVar(&kubeconfig, "kubeconfig", "", "Path to the kubeconfig file")
	monitorCmd.MarkFlagRequired("namespace")
}
