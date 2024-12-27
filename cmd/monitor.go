package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/ethanhorrigan/kubeskope/pkg/analysis"
	"github.com/ethanhorrigan/kubeskope/pkg/k8s"
	"github.com/ethanhorrigan/kubeskope/pkg/output"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/spf13/cobra"
)

var namespace string
var kubeconfig string

type model struct {
	rows       [][]string
	loading    bool
	error      error
	namespace  string
	kubeconfig string
}

type tickMsg struct{}
type refreshMsg struct{
	rows [][]string
	error error
}

func NewModel(namespace, kubeconfig string) model {
	return model{
		namespace:  namespace,
		kubeconfig: kubeconfig,
		loading:    true,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(fetchMetrics(m.namespace, m.kubeconfig), tickCmd())
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	case tickMsg:
		return m, fetchMetrics(m.namespace, m.kubeconfig)
	case refreshMsg:
		m.rows = msg.rows
		m.error = msg.error
		m.loading = false
	}
	return m, tickCmd()
}

func (m model) View() string {
	if m.loading {
		return "Loading metrics...\n"
	}

	if m.error != nil {
		return fmt.Sprintf("Error: %v\n", m.error)
	}

	return output.RenderBubbleTable(m.rows)
}

func tickCmd() tea.Cmd {
	return tea.Tick(2*time.Second, func(time.Time) tea.Msg {
		return tickMsg{}
	})
}

func fetchMetrics(namespace, kubeconfig string) tea.Cmd {
	return func() tea.Msg {
		client, config, err := k8s.GetKubeClient(kubeconfig)
		if err != nil {
			return refreshMsg{rows: nil, error: err}
		}

		pods, err := client.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
		if err != nil {
			return refreshMsg{rows: nil, error: err}
		}

		usageMetrics, err := k8s.GetPodMetrics(config, namespace)
		if err != nil {
			return refreshMsg{rows: nil, error: err}
		}

		var rows [][]string
		threshold := 90.0

		for _, pod := range pods.Items {
			podRows := analysis.AnalyzePod(pod, usageMetrics, threshold)
			rows = append(rows, podRows...)
		}

		return refreshMsg{rows: rows, error: nil}
	}
}

func ExecuteMonitor() {
	p := tea.NewProgram(NewModel(namespace, kubeconfig))
	if err := p.Start(); err != nil {
		log.Fatalf("Error running application: %v", err)
		os.Exit(1)
	}
}

func init() {
	monitorCmd := &cobra.Command{
		Use:   "monitor",
		Short: "Monitor resource usage in a namespace",
		Run: func(cmd *cobra.Command, args []string) {
			ExecuteMonitor()
		},
	}

	// Attach flags to the monitorCmd
	monitorCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "Namespace to monitor")
	monitorCmd.Flags().StringVar(&kubeconfig, "kubeconfig", "", "Path to the kubeconfig file")
	monitorCmd.MarkFlagRequired("namespace")

	// Add monitorCmd to rootCmd
	rootCmd.AddCommand(monitorCmd)
}

