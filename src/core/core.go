package core

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"okinawa/log"
)

var window *glfw.Window
var program uint32

// Initialize prepares the core component of the engine (handles OpenGL)
func Initialize(windowWidth int, windowHeight int) {
	window = initGlfw(windowWidth, windowHeight)
	defer glfw.Terminate()

	program = initOpenGL()

	for !window.ShouldClose() {
		draw(window, program)
	}
}

// initGlfw initializes glfw and returns a Window to use
func initGlfw(windowWidth int, windowHeight int) *glfw.Window {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(windowWidth, windowHeight, "okinawa.go", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()

	return window
}

// initOpenGL initializes OpenGL and returns an intiialized program
func initOpenGL() uint32 {
	err := gl.Init()
	if err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Log("core::initOpenGL", "OpenGL version"+version)

	program := gl.CreateProgram()
	gl.LinkProgram(program)

	return program
}

func draw(window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	glfw.PollEvents()
	window.SwapBuffers()
}
