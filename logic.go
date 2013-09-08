package main

import (
	al "github.com/tapir/allegro5"
)

const (
	dt           = 0.03
	maxFrameTime = 0.2
)

func processLogic(alphaOut chan<- float64, inputIn <-chan uint) {
	var (
		t           = 0.0
		currentTime = al.GetTime()
		accumulator = 0.0
	)

	// Main logic loop
	for {
		newTime := al.GetTime()
		frameTime := newTime - currentTime
		currentTime = newTime

		if frameTime > maxFrameTime {
			frameTime = maxFrameTime
		}

		accumulator += frameTime

		for accumulator >= dt {
			<-inputIn
			// Do stuff here
			t += dt
			accumulator -= dt
		}

		// Switch to render coroutine by sending the interpolation
		alphaOut <- accumulator / dt
	}
}
