package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dev-provision",
	Short: "DevProGo is a Go-based application that empowers developers to swiftly provision development environments",
	Long: `DevProGo is a Go-based application that empowers developers to swiftly provision development environments 
resembling production setups using Infrastructure as Code (IaC) principles.
Simplify and accelerate the process of creating development environments with ease.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
