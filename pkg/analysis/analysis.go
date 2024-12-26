package analysis

import (
	// "fmt"
	corev1 "k8s.io/api/core/v1"
)

// CalculateUtilization calculates CPU/Memory utilization percentages
func CalculateUtilization(requested, used int64) float64 {
	if requested == 0 {
		return 0
	}
	return (float64(used) / float64(requested)) * 100
}

// AnalyzePod analyzes a pod's resource usage
func AnalyzePod(pod corev1.Pod) [][]string {
	var results [][]string

	for _, container := range pod.Spec.Containers {
		results = append(results, []string{
			pod.Name,
			container.Name,
			container.Resources.Requests.Cpu().String(),
			container.Resources.Requests.Memory().String(),
			container.Resources.Limits.Cpu().String(),
			container.Resources.Limits.Memory().String(),
		})
	}

	return results
}
