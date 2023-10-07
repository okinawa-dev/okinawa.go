package item

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var (
	spriteName = "Item"
	visible    = true
	frame      = []float32{
		0.5, 0.5, 0, // top right
		-0.5, 0.5, 0, // top left
		-0.5, -0.5, 0, // bottom left
		0.5, -0.5, 0, // bottom right
	}
	// []float32{
	// 	0, 0.5, 0, // top
	// 	-0.5, -0.5, 0, // left
	// 	0.5, -0.5, 0, // right
	// }
)

// Initialize prepares an item
func Initialize() {
	var vao uint32

	vao = makeVertexArrayObject(frame)
	gl.BindVertexArray(vao)
}

// GetVisible returns the value of the visible property
func GetVisible() bool {
	return visible
}

// SetVisible sets the value of the visible property
func SetVisible(value bool) {
	visible = value
}

// Draw renders the item
func Draw(window *glfw.Window, program uint32) {
	// use gl.TRIANGLES to draw triangles using vertices in groups of three
	// use gl.TRIANGLE_FAN to draw the first 3, then 2,3,4, then 3,4,5, etc.
	gl.DrawArrays(gl.TRIANGLE_FAN, 0, int32(len(frame)/3))
}

// makeVertexArrayObject initializes and returns a vertex array from the points provided
func makeVertexArrayObject(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}
