package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kubeskope",
	Short: "KubeSkope is a Kubernetes resource monitor",
	Long:  `KubeSkope provides insights into Kubernetes resource usage and optimization.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
