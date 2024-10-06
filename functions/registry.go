package functions

import (
	"fmt"
	"golang.org/x/sys/windows/registry"
	"log"
)

// DisableTelemetry Disables \DataCollection through registry modification.
func DisableTelemetry() {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Policies\Microsoft\Windows\DataCollection`,
		registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		log.Println("error while Opening Registry Key, Error: ", err)
	}
	defer func(key registry.Key) {
		err := key.Close()
		if err != nil {
			log.Println("error while closing the deferred open registry key, error: ", err)
		}
	}(k)

	errors := k.SetDWordValue("AllowTelemetry", 0)
	if errors != nil {
		log.Println("error while setting DWord value, error: ", errors)
	} else {
		value, _, err := k.GetIntegerValue("AllowTelemetry")
		if err != nil {
			log.Println("error getting AllowTelemetry value:", err)
		} else {
			fmt.Printf("AllowTelemetry current value: %d\n", value)
		}
	}
}

// DisableWeatherAndNews Disables annoying weather and news popups in start-menu and lockscreen.
func DisableWeatherAndNews() {
	keyPath := `SOFTWARE\Policies\Microsoft\Dsh`
	k, _, err := registry.CreateKey(registry.LOCAL_MACHINE, keyPath, registry.SET_VALUE)
	if err != nil {
		log.Println("error while creating registry key, error: ", err)
	}
	defer func(registry registry.Key) {
		err := registry.Close()
		if err != nil {
			log.Println("error while closing the deferred open registry key, error: ", err)
		}
	}(k)

	newDWord := k.SetDWordValue("AllowNewsAndInterests", 0)
	if newDWord != nil {
		log.Println("error while setting DWord Value, error: ", newDWord)
	}
}

// RemoveBing Disables Bing search in start menu search promoting explorer local search.
func RemoveBing() {
	keyPath := `Software\Policies\Microsoft\Windows\Explorer`
	k, _, err := registry.CreateKey(registry.LOCAL_MACHINE, keyPath, registry.SET_VALUE)
	if err != nil {
		log.Println("error while creating registry key, error: ", err)
	}
	defer func(registry registry.Key) {
		err := registry.Close()
		if err != nil {
			log.Println("error while closing the deferred open registry key, error: ", err)
		}
	}(k)

	newDWord := k.SetDWordValue("DisableSearchBoxSuggestions", 0)
	if newDWord != nil {
		log.Println("error while closing the deferred open registry key, error: ", err)
	}
}

// DisableCopilot Disables Copilot through registry edits.
func DisableCopilot() {
	keyPath := `SOFTWARE\Policies\Microsoft\Windows\WindowsCopilot`
	k, _, err := registry.CreateKey(registry.LOCAL_MACHINE, keyPath, registry.SET_VALUE)
	if err != nil {
		log.Println("error while creating registry key, error: ", err)
	}
	defer func(key registry.Key) {
		err := key.Close()
		if err != nil {
			log.Println("error while closing the deferred open registry key, error: ", err)
		}
	}(k)

	newDWord := k.SetDWordValue("TurnOffWindowsCopilot", 1)
	if newDWord != nil {
		log.Println("error while setting DWord value, error:", newDWord)
	}
}

// DisableAdsExplorer Disables Ads in explorer.
func DisableAdsExplorer() {
	keyPath := `SOFTWARE\Microsoft\Windows\CurrentVersion\Explorer\Advanced`
	k, _, err := registry.CreateKey(registry.CURRENT_USER, keyPath, registry.SET_VALUE)
	if err != nil {
		log.Println("error while creating registry key, error: ", err)
	}
	defer func(key registry.Key) {
		err := key.Close()
		if err != nil {
			log.Println("error while closing the deferred open registry key, error: ", err)
		}
	}(k)
	newDWord := k.SetDWordValue("ShowSyncProviderNotifications", 0)
	if newDWord != nil {
		log.Println("error while setting DWord Value, error: ", newDWord)
	}
}
