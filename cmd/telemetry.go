package cmd

import (
	WinDebloat "WinDebloat-Go/functions"
	"github.com/spf13/cobra"
)

var disableTelemetryCmd = &cobra.Command{
	Use:   "disable-telemetry",
	Short: "Disables Windows telemetry",
	Run: func(cmd *cobra.Command, args []string) {
		WinDebloat.DisableTelemetry()
	},
}
