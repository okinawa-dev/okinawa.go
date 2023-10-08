package item

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"okinawa/math"
)

// Item defines a drawable object
type Item struct {
	spriteName string
	visible    bool

	position math.Point3
	size     math.Point2
}

var (
	frame = []float32{
		0.5, 0.5, 0, // top right
		-0.5, 0.5, 0, // top left
		-0.5, -0.5, 0, // bottom left
		0.5, -0.5, 0, // bottom right
	}
)

// Initialize prepares an item
func (i Item) Initialize() {
	i.spriteName = "Item"
	i.visible = true
	i.position = math.Point3{X: 0, Y: 0, Z: 0}
	i.size = math.Point2{X: 0.5, Y: 0.5}

	makeVertexArrayObject(frame)
}

// GetVisible returns the value of the visible property
func (i Item) GetVisible() bool {
	return i.visible
}

// SetVisible sets the value of the visible property
func (i Item) SetVisible(value bool) {
	i.visible = value
}

// GetPosition returns the position of the item
func (i Item) GetPosition() math.Point3 {
	return i.position
}

// SetPosition sets the position of the item
func (i Item) SetPosition(value math.Point3) {
	i.position = value
}

// GetSize returns the position of the item
func (i Item) GetSize() math.Point2 {
	return i.size
}

// SetSize sets the position of the item
func (i Item) SetSize(value math.Point2) {
	i.size = value
}

// Draw renders the item
func (i Item) Draw(window *glfw.Window, program uint32) {
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
	// gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}
