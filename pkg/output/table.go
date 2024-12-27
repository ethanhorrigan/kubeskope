package output

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

// RenderTable renders resource usage data in a table
func RenderTable(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Pod", "Container", "CPU Requests", "Memory Requests", "CPU Limits", "Memory Limits", "CPU Usage", "Memory Usage"})

	table.SetRowLine(true)

	for _, row := range data {
		table.Append(row)
	}
	table.Render()
}
