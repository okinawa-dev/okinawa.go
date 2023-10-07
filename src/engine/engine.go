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

func Initialize() {
	log.Log("engine::Initialize", "Let's go!")

	runtime.LockOSThread()

	log.Initialize()
	core.Initialize(windowWidth, windowHeight)
}
