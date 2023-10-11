package item

import (
	// "fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"

	"okinawa/log"
	"okinawa/math"
)

// Item defines a drawable object
type Item struct {
	name    string
	visible bool

	position math.Point2
	size     math.Point2
	rotation math.Rotation

	// item hierarchy
	parent        *Item
	attachedItems []*Item

	// openGL information
	vertexArrayObject  uint32
	vertexBufferObject uint32
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
	i.name = "Item"
	i.visible = true
	i.position = math.Point2{X: 0, Y: 0}
	i.size = math.Point2{X: 0.5, Y: 0.5}

	i.parent = nil
	i.attachedItems = []*Item{}

	log.Info("item::Initialize", i.name)
	i.initializeVertexArrayObject(frame)
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
func (i Item) GetPosition() math.Point2 {
	// return i.position
	var (
		result              math.Point2
		parentPosition      math.Point2
		transformedPosition math.Point2
	)

	if i.GetParent() != nil {
		parentPosition = i.GetParent().GetPosition()
		transformedPosition = i.GetParent().GetRotation().TransformPosition(i.GetPosition())

		result.X = parentPosition.X + transformedPosition.X
		result.Y = parentPosition.Y + transformedPosition.Y
	} else {
		result = i.position
	}

	return result
}

// SetPosition sets the position of the item
func (i Item) SetPosition(value math.Point2) {
	i.position = value
}

// GetRotation returns the rotation of the item
func (i Item) GetRotation() math.Rotation {
	return i.rotation
}

// GetSize returns the position of the item
func (i Item) GetSize() math.Point2 {
	return i.size
}

// SetSize sets the position of the item
func (i Item) SetSize(value math.Point2) {
	i.size = value
}

// GetParent returns the parent of the item
func (i Item) GetParent() *Item {
	return i.parent
}

// SetParent sets the parent of the item
func (i Item) SetParent(value *Item) {
	i.parent = value
}

// Draw renders the item
func (i Item) Draw(window *glfw.Window, program uint32) {

	// TODO check
	// update the pipeline state with the buffer we intend to use
	gl.BindBuffer(gl.ARRAY_BUFFER, i.vertexBufferObject)

	// use gl.TRIANGLES to draw triangles using vertices in groups of three
	// use gl.TRIANGLE_FAN to draw the first 3, then 2,3,4, then 3,4,5, etc.
	gl.DrawArrays(gl.TRIANGLE_FAN, 0, int32(len(frame)/3))
}

// initializeVertexArrayObject initializes and returns a vertex array from the points provided
func (i Item) initializeVertexArrayObject(points []float32) {
	// generate 1 buffer object
	gl.GenBuffers(1, &i.vertexBufferObject)
	// vbo will contain an array of vertices
	gl.BindBuffer(gl.ARRAY_BUFFER, i.vertexBufferObject)
	// fill the buffer with data
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	log.Info("item::initializeVertexArrayObject" /* fmt.Sprint(i.name) */, i.GetPosition().String())

	gl.GenVertexArrays(1, &i.vertexArrayObject)
	gl.BindVertexArray(i.vertexArrayObject)
	// vertex position is treated as vertex attribute index 0 in the fixed function pipeline
	gl.EnableVertexAttribArray(0)
	// tell the pipeline how to interpret the data inside the buffer
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
}
