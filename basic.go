package main

import (
	"github.com/wailsapp/wails/v2"
)

// Basic application struct
type Galaxy struct {
	Name    string
	runtime *wails.Runtime
}

// NewGalaxy
func NewGalaxy(name string) *Galaxy {
	return &Galaxy{
		Name: name,
	}
}

// startup is called at application startup
func (g *Galaxy) startup(runtime *wails.Runtime) {
	// Perform your setup here
	g.runtime = runtime
	runtime.Window.SetTitle(g.Name)
}

// shutdown is called at application termination
func (g *Galaxy) shutdown() {
	// Perform your teardown here
}

// Speed
func (g *Galaxy) Speed() string {
	return "1.3 million"
}

// PlanetSpeed
func (g *Galaxy) PlanetSpeed(planet string) string {
	switch planet {
	case "mercury":
		return "107,082"
	case "venus":
		return "78,337"
	case "earth":
		return "67,000"
	case "mars":
		return "53,853"
	case "jupiter":
		return "29,236"
	case "saturn":
		return "21,675"
	case "uranus":
		return "15,233"
	case "neptune":
		return "12,146"

	// fall back to earth
	default:
		return "67,000"
	}
}
