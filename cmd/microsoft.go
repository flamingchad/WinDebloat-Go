package cmd

import (
	"WinDebloat-Go/functions"
	"github.com/spf13/cobra"
)

var disableMicrosoftTelemetryCmd = &cobra.Command{
	Use:   "disable-microsoft-telemetry",
	Short: "Blocks Microsoft telemetry in the hosts file",
	Run: func(cmd *cobra.Command, args []string) {
		functions.DisableMicrosoftTelemetry()
	},
}

var disableWeatherAndNewsCmd = &cobra.Command{
	Use:   "disable-weather-and-news",
	Short: "Blocks Weather and News",
	Run: func(cmd *cobra.Command, args []string) {
		functions.DisableWeatherAndNews()
	},
}

var removeBing = &cobra.Command{
	Use:   "remove-bing",
	Short: "Removes Bing",
	Run: func(cmd *cobra.Command, args []string) {
		functions.RemoveBing()
	},
}

var disableCopilot = &cobra.Command{
	Use:   "disable-copliot",
	Short: "Removes Copliot",
	Run: func(cmd *cobra.Command, args []string) {
		functions.DisableCopilot()
	},
}
