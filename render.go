package main

import (
	al "github.com/tapir/allegro5"
	"runtime"
	"log"
)

// Some settings
const (
	title  = "Core Game"
	width  = 800
	height = 600
)

func render(alphaIn <-chan float64, displayOut chan<- *al.Display) {
	// Lock the render thread
	runtime.LockOSThread()
	
	// Setup display flag and options
	al.SetNewDisplayFlags(al.Windowed | al.Opengl30)
	al.SetNewDisplayOption(al.Vsync, 1, al.Suggest)
	al.SetNewDisplayOption(al.DepthSize, 0, al.Suggest)
	al.SetNewDisplayOption(al.SampleBuffers, 1, al.Suggest)
	al.SetNewDisplayOption(al.UpdateDisplayRegion_, 1, al.Suggest)
	al.SetNewDisplayOption(al.RedSize, 8, al.Require)
	al.SetNewDisplayOption(al.GreenSize, 8, al.Require)
	al.SetNewDisplayOption(al.BlueSize, 8, al.Require)
	al.SetNewDisplayOption(al.AlphaSize, 8, al.Require)

	// Create window
	display := al.CreateDisplay(width, height)
	if display == nil {
		log.Fatalln("Could not create display.")
	}
	defer display.Destroy()
	
	// Set title
	display.SetWindowTitle(title)
	
	// Send display to main goroutine
	displayOut <- display
	
	// Make context current in this goroutine
	display.SetTargetBackbuffer()
	
	// Main rendering loop
	for {
		//Do things with interpolation value
		<-alphaIn
		al.ClearToColor(al.MapRgb(255, 0, 255))

		al.FlipDisplay()
	}
}
