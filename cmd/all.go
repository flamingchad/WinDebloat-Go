package cmd

import (
	"WinDebloat-Go/functions"
	"github.com/spf13/cobra"
)

var all = &cobra.Command{
	Use:   "all",
	Short: "Applies all the recommended settings",
	Run: func(cmd *cobra.Command, args []string) {
		functions.DisableDiagTrack()
		functions.DisableMicrosoftTelemetry()
		functions.RemoveBing()
		functions.DisableCopilot()
		functions.DisableTelemetry()
	},
}
