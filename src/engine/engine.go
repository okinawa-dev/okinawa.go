package engine

import (
	"log"
	"runtime"

	"okinawa/core"
)

const (
	windowWidth  = 640
	windowHeight = 480
)

func Initialize() {
	log.Println("Let's begin!")

	runtime.LockOSThread()

	core.Initialize(windowWidth, windowHeight)
}
