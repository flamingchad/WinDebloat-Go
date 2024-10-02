package main

import (
	"github.com/spf13/cobra"
)

var all = &cobra.Command{
	Use:   "all",
	Short: "Applies all the recommended settings",
	Run: func(cmd *cobra.Command, args []string) {
		DisableDiagTrack()
		DisableMicrosoftTelemetry()
		RemoveBing()
		DisableCopilot()
		DisableTelemetry()
	},
}
