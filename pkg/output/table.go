package output

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	headers = []string{"POD", "CONTAINER", "CPU REQUESTS", "MEMORY REQUESTS", "CPU LIMITS", "MEMORY LIMITS", "CPU USAGE", "MEMORY USAGE"}
	columnWidths = []int{30, 20, 15, 15, 15, 15, 15, 15} // Fixed column widths
	headersStyle = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205"))
	rowStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("245"))
	selectedRowStyle = lipgloss.NewStyle().
		Background(lipgloss.Color("236")).
		Foreground(lipgloss.Color("229"))
)

// padString ensures strings fit a fixed width
func padString(input string, width int) string {
	if len(input) > width {
		return input[:width-3] + "..." // Truncate and add ellipsis if too long
	}
	return fmt.Sprintf("%-*s", width, input) // Pad with spaces
}

// RenderBubbleTable renders a table for Bubble Tea-based UI
func RenderBubbleTable(rows [][]string) string {
	var b strings.Builder

	// Render headers with fixed widths
	for i, header := range headers {
		b.WriteString(headersStyle.Render(padString(header, columnWidths[i])))
	}
	b.WriteString("\n")

	// Render rows
	for i, row := range rows {
		for j, cell := range row {
		    cell = padString(cell, columnWidths[j])
			if i%2 == 0 {
				b.WriteString(selectedRowStyle.Render(cell))
			} else {
				b.WriteString(rowStyle.Render(cell))
			}
		}
		b.WriteString("\n")
	}

	return b.String()
}
