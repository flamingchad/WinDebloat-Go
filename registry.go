package main

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"log"
)

func disableTelemetry() {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Policies\Microsoft\Windows\DataCollection`,
		registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		log.Println(err)
	}
	defer func(key registry.Key) {
		err := key.Close()
		if err != nil {
			log.Println(err)
		}
	}(k)

	errors := k.SetDWordValue("AllowTelemetry", 0)
	if errors != nil {
		log.Println(errors)
	} else {
		value, _, err := k.GetIntegerValue("AllowTelemetry")
		if err != nil {
			log.Println("Error getting AllowTelemetry value:", err)
		} else {
			fmt.Printf("AllowTelemetry current value: %d\n", value)
		}
	}
}
