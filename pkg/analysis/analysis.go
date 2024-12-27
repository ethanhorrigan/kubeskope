package analysis

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"strings"
)

// CalculateUtilization calculates CPU/Memory utilization percentages
func CalculateUtilization(requested, used int64) float64 {
	if requested == 0 {
		return 0
	}
	return (float64(used) / float64(requested)) * 100
}

func GenerateBar(percentage float64, threshold float64) string {
	barLength := 20
	filledLength := int(percentage / 100 * float64(barLength))
	bar := strings.Repeat("█", filledLength) + strings.Repeat("░", barLength-filledLength)

	if percentage > threshold {
		return fmt.Sprintf("\033[1;31m%s\033[0m (%.2f%%)", bar, percentage) // Red for over-threshold
	}

	return fmt.Sprintf("\033[1;32m%s\033[0m (%.2f%%)", bar, percentage) // Green otherwise
}

// AnalyzePod analyzes a pod's resource usage
func AnalyzePod(pod corev1.Pod, usage map[string]map[string]string, threshold float64) [][]string {
	var results [][]string

	for _, container := range pod.Spec.Containers {
		usageValues := usage[pod.Name][container.Name]
		var cpuUsage, memUsage float64
		fmt.Sscanf(usageValues, "CPU: %fm, Memory: %fMi", &cpuUsage, &memUsage)

		// Calculate thresholds
		cpuRequest := container.Resources.Requests.Cpu().MilliValue()
		memRequest := container.Resources.Requests.Memory().Value() / 1024 / 1024

		cpuPercentage := CalculateUtilization(cpuRequest, int64(cpuUsage))
		memPercentage := CalculateUtilization(memRequest, int64(memUsage))

		// Generate bars
		cpuBar := GenerateBar(cpuPercentage, threshold)
		memBar := GenerateBar(memPercentage, threshold)

		results = append(results, []string{
			pod.Name,
			container.Name,
			container.Resources.Requests.Cpu().String(),
			container.Resources.Requests.Memory().String(),
			container.Resources.Limits.Cpu().String(),
			container.Resources.Limits.Memory().String(),
			cpuBar,
			memBar,
		})
	}

	return results
}
