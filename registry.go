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

func disableWeatherAndNews() {
	keyPath := `SOFTWARE\Policies\Microsoft\Dsh`
	k, _, err := registry.CreateKey(registry.LOCAL_MACHINE, keyPath, registry.SET_VALUE)
	if err != nil {
		log.Println(err)
	}
	defer func(registry registry.Key) {
		err := registry.Close()
		if err != nil {
			log.Println(err)
		}
	}(k)

	newDWord := k.SetDWordValue("AllowNewsAndInterests", 0)
	if newDWord != nil {
		log.Println(newDWord)
	}
}

func removeBing() {
	keyPath := `Software\Policies\Microsoft\Windows\Explorer`
	k, _, err := registry.CreateKey(registry.LOCAL_MACHINE, keyPath, registry.SET_VALUE)
	if err != nil {
		log.Println(err)
	}
	defer func(registry registry.Key) {
		err := registry.Close()
		if err != nil {
			log.Println(err)
		}
	}(k)

	newDWord := k.SetDWordValue("DisableSearchBoxSuggestions", 0)
	if newDWord != nil {
		log.Println(newDWord)
	}
}

func disableCopilot() {
	keyPath := `SOFTWARE\Policies\Microsoft\Windows\WindowsCopilot`
	k, _, err := registry.CreateKey(registry.LOCAL_MACHINE, keyPath, registry.SET_VALUE)
	if err != nil {
		log.Println(err)
	}
	defer func(key registry.Key) {
		err := key.Close()
		if err != nil {
			log.Println(err)
		}
	}(k)

	newDWord := k.SetDWordValue("TurnOffWindowsCopilot", 1)
	if newDWord != nil {
		log.Println(newDWord)
	}
}
