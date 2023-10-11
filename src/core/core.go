package core

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"okinawa/item"
	"okinawa/log"
	"okinawa/math"
	"okinawa/utils"
)

var (
	window  *glfw.Window
	program uint32

	// TODO
	i1 item.Item
	i2 item.Item
)

// Initialize prepares the core component of the engine (handles OpenGL)
func Initialize(windowWidth int, windowHeight int) {

	window = initGlfw(windowWidth, windowHeight)
	defer glfw.Terminate()

	program = initOpenGL()

	// after opengl is initialized
	i1 = item.Item{Name: "Item1"}
	i1.Initialize()
	i1.SetPosition(&math.Point2{X: 0, Y: 0})
	i1.SetSize(&math.Point2{X: 20, Y: 20})
	i1.SetRotation(10)

	i2 = item.Item{Name: "Item2"}
	i2.Initialize()
	i2.SetPosition(&math.Point2{X: 10, Y: 10})
	i2.SetSize(&math.Point2{X: 20, Y: 20})
	i2.SetRotation(10)

	window.SetInputMode(glfw.StickyKeysMode, glfw.True)

	for window.GetKey(glfw.KeyEscape) != glfw.Press && !window.ShouldClose() {
		Draw(window, program)
	}
}

// initGlfw initializes glfw and returns a Window to use
func initGlfw(windowWidth int, windowHeight int) *glfw.Window {
	err := glfw.Init()
	if err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Samples, 4) // 4x antialiasing
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // We want OpenGL 4.1
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile) // We don't want the old OpenGL
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)    // To make MacOS happy; should not be needed

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
	log.Info("core::initOpenGL", "OpenGL version"+version)

	// read basic vertex shader from file
	vertexShaderSource, err := utils.GetFileContents("shaders/simplevertexshader.glsl")
	if err != nil {
		panic(err)
	}

	log.Info("core::initOpenGL", "vertex shader:\n"+vertexShaderSource)

	vertexShader, err := compileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	// read basic fragment shader from file
	fragmentShaderSource, err := utils.GetFileContents("shaders/simplefragmentshader.glsl")
	if err != nil {
		panic(err)
	}

	log.Info("core::initOpenGL", "fragment shader:\n"+fragmentShaderSource)

	fragmentShader, err := compileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	program := gl.CreateProgram()
	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	return program
}

func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("Failed to compile %v: %v", source, log)
	}

	return shader, nil
}

// Draw renders the window
func Draw(window *glfw.Window, program uint32) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(program)

	// TODO draw the full scene
	i1.Draw(window, program)
	i2.Draw(window, program)

	window.SwapBuffers()
	glfw.PollEvents()
}
