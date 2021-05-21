package main

import (
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

func main() {

	// Create application with options
	galaxy := NewGalaxy("The Milky Way")

	err := wails.Run(&options.App{
		Title:         "The Galaxy",
		Width:         740,
		Height:        569,
		DisableResize: true,
		Mac: &mac.Options{
			WebviewIsTransparent:          true,
			WindowBackgroundIsTranslucent: true,
			TitleBar:                      mac.TitleBarDefault(),
			Menu:                          menu.DefaultMacMenu(),
		},
		LogLevel: logger.DEBUG,
		Startup:  galaxy.startup,
		Shutdown: galaxy.shutdown,
		Bind: []interface{}{
			galaxy,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}
