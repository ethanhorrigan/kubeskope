package analysis

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
)

// CalculateUtilization calculates CPU/Memory utilization percentages
func CalculateUtilization(requested, used int64) float64 {
	if requested == 0 {
		return 0
	}
	return (float64(used) / float64(requested)) * 100
}

func padString(input string, width int) string {
	if len(input) > width {
		return input[:width-3] + "..." // Truncate and add ellipsis if too long
	}
	return fmt.Sprintf("%-*s", width, input) // Pad with spaces
}

// AnalyzePod analyzes a pod's resource usage
func AnalyzePod(pod corev1.Pod, usage map[string]map[string]string, threshold float64) [][]string {
	var results [][]string

	for _, container := range pod.Spec.Containers {
		usageValues := usage[pod.Name][container.Name]
		var cpuUsage, memUsage float64
		_, err := fmt.Sscanf(usageValues, "CPU: %fm, Memory: %fMi", &cpuUsage, &memUsage)
		if err != nil {
			cpuUsage, memUsage = 0, 0 // Handle missing or malformed data
		}

		// Calculate thresholds
		cpuRequest := container.Resources.Requests.Cpu().MilliValue()
		memRequest := container.Resources.Requests.Memory().Value() / 1024 / 1024

		cpuPercentage := CalculateUtilization(cpuRequest, int64(cpuUsage))
		memPercentage := CalculateUtilization(memRequest, int64(memUsage))

		results = append(results, []string{
			padString(pod.Name, 30),
			padString(container.Name, 20),
			padString(container.Resources.Requests.Cpu().String(), 15),
			padString(container.Resources.Requests.Memory().String(), 15),
			padString(container.Resources.Limits.Cpu().String(), 15),
			padString(container.Resources.Limits.Memory().String(), 15),
			fmt.Sprintf("%.2f%%", cpuPercentage),
			fmt.Sprintf("%.2f%%", memPercentage),
		})
	}

	return results
}
