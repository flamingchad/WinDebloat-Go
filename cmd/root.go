package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "debloat",
	Short: "A tool to debloat Windows 11",
	Long:  `This tool disables telemetry, tracking, and unnecessary services from Windows 11.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run 'debloat --help' to see available commands.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Add subcommands here
	rootCmd.AddCommand(all)
	rootCmd.AddCommand(disableDiagTrackCmd)
	rootCmd.AddCommand(disableTelemetryCmd)
	rootCmd.AddCommand(disableMicrosoftTelemetryCmd)
	rootCmd.AddCommand(disableWeatherAndNewsCmd)
	rootCmd.AddCommand(disableCopilot)
	rootCmd.AddCommand(removeBing)
}
