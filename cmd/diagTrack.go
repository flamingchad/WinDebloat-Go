package cmd

import (
	"WinDebloat-Go/functions"
	"github.com/spf13/cobra"
)

var disableDiagTrackCmd = &cobra.Command{
	Use:   "disable-diagtrack",
	Short: "Disables the DiagTrack service",
	Run: func(cmd *cobra.Command, args []string) {
		functions.DisableDiagTrack()
	},
}
