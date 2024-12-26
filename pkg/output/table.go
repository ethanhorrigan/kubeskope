package output

import (
    "os"
    "github.com/olekukonko/tablewriter"
)

// RenderTable renders resource usage data in a table
func RenderTable(data [][]string) {
    table := tablewriter.NewWriter(os.Stdout)
    table.SetHeader([]string{"Pod", "Container", "CPU Requests", "Memory Requests", "CPU Limits", "Memory Limits"})
    table.SetRowLine(true)

    for _, row := range data {
        table.Append(row)
    }
    table.Render()
}
