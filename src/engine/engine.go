package engine

import (
	"runtime"

	"okinawa/core"
	"okinawa/log"
)

const (
	windowWidth  = 640
	windowHeight = 480
)

// Initialize prepares the engine, and calls Initialize in every engine component
func Initialize() {
	log.Info("engine::Initialize", "Let's go!")

	runtime.LockOSThread()

	log.Initialize()
	core.Initialize(windowWidth, windowHeight)
}
