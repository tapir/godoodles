package main

import (
	al "github.com/tapir/allegro5"
	"github.com/tapir/allegro5/imageio"
	"runtime"
	"log"
)

func init() {
	// Lock the main thread
	runtime.LockOSThread()
}

func main() {
	// Prepare channels
	alpha := make(chan float64) // Logic-render interpolation value
	input := make(chan uint) // Input data
	display := make(chan *al.Display) // Display created

	// Start Allegro
	if !al.Init() {
		log.Fatalln("Could not start Allegro.")
	}
	//defer al.UninstallSystem()

	// Load image addon
	if !imageio.Init() {
		log.Fatalln("Could not start image IO addon.")
	}
	imageio.Shutdown()

	// Create an event queue
	queue := al.CreateEventQueue()
	if queue == nil {
		log.Fatalln("Could not create event queue.")
	}
	defer queue.Destroy()

	// Process logic
	go processLogic(alpha, input)
	
	// Process render
	go render(alpha, display)

	// Get display created in render goroutine
	d := <-display
	
	// Register event sources
	queue.RegisterEventSource(d.GetEventSource())

	// Start timer and process input
	done := false
	for !done {
		ok, event := queue.GetNextEvent()
		if ok {
			// Handle all allegro events here
			if event.Type == al.EventDisplayClose {
				done = true
			}
		}
		
		// Switch context to logic goroutine by sending input data
		input <- 0
	}
}
