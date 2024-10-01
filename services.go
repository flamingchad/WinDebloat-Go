package main

import (
	"fmt"
	"os/exec"
)

// Goal: Implement the function to stop and disable Windows telemetry services like DiagTrack.
// Steps:
//
//	Write a function that uses os/exec to run PowerShell commands like Stop-Service or Set-Service -StartupType Disabled.
//	Use error handling to deal with cases where the service is already stopped or doesn't exist.

func disableDiagTrack() {
	command := `Start-Process powershell -ArgumentList 'Stop-Service DiagTrack; Set-Service -Name DiagTrack -StartupType Disabled' -Verb RunAs`
	out, err := exec.Command("powershell.exe", "-Command", command).Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Service stopped successfully.")
	}
	fmt.Println(string(out))
}
