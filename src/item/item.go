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
	Name    string
	visible bool

	position *math.Point2 // position on screen (relative to parent)
	size     *math.Point2 // size in pixels
	rotation *math.Rotation

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
func (i *Item) Initialize() {
	if i.Name == "" {
		i.Name = "something"
	}

	i.visible = true
	i.position = &math.Point2{X: 10, Y: 110}
	i.size = &math.Point2{X: 10, Y: 10}
	// zero degrees, and matrix for zero degrees rotation
	i.rotation = &math.Rotation{Angle: 0, A: 1, B: 0, C: 0, D: 1}

	i.parent = nil
	i.attachedItems = []*Item{}

	log.Info("item::Initialize", i.Name, *i.rotation)
	i.initializeVertexArrayObject(frame)
}

// GetVisible returns the value of the visible property
func (i *Item) GetVisible() bool {
	return i.visible
}

// SetVisible sets the value of the visible property
func (i *Item) SetVisible(value bool) {
	i.visible = value
}

// GetPosition returns the global position of the item on screen
func (i *Item) GetPosition() *math.Point2 {
	// return i.position
	var (
		result              *math.Point2
		parentPosition      *math.Point2
		transformedPosition *math.Point2
	)

	if i.parent != nil {
		parentPosition = i.parent.GetPosition()
		transformedPosition = i.parent.GetRotation().TransformPosition(i.position)

		result.X = parentPosition.X + transformedPosition.X
		result.Y = parentPosition.Y + transformedPosition.Y
	} else {
		result = i.position
	}

	return result
}

// SetPosition sets the position of the item
func (i *Item) SetPosition(value *math.Point2) {
	i.position = value
}

// GetRotation returns the global rotation object for the item
func (i *Item) GetRotation() *math.Rotation {
	if i.parent != nil {
		i.rotation.Add(i.parent.GetRotation())
	}

	return i.rotation
}

// GetRotationValue returns the global rotation of the item
func (i *Item) GetRotationValue() float64 {
	if i.parent != nil {
		return i.rotation.GetAngle() + i.parent.GetRotationValue()
	}

	return i.rotation.GetAngle()
}

// SetRotation sets the rotation of the item
func (i *Item) SetRotation(value float64) {
	i.rotation.Update(value)
}

// GetSize returns the position of the item
func (i *Item) GetSize() *math.Point2 {
	return i.size
}

// SetSize sets the position of the item
func (i *Item) SetSize(value *math.Point2) {
	i.size = value
}

// GetParent returns the parent of the item
func (i *Item) GetParent() Item {
	return *i.parent
}

// SetParent sets the parent of the item
func (i *Item) SetParent(value *Item) {
	i.parent = value
}

// Draw renders the item
func (i *Item) Draw(window *glfw.Window, program uint32) {

	log.Info("item::Draw", i.Name, "position", i.GetPosition(), "rotation", i.GetRotation())

	// TODO check
	// update the pipeline state with the buffer we intend to use
	gl.BindBuffer(gl.ARRAY_BUFFER, i.vertexBufferObject)

	// use gl.TRIANGLES to draw triangles using vertices in groups of three
	// use gl.TRIANGLE_FAN to draw the first 3, then 2,3,4, then 3,4,5, etc.
	gl.DrawArrays(gl.TRIANGLE_FAN, 0, int32(len(frame)/3))
}

// initializeVertexArrayObject initializes and returns a vertex array from the points provided
func (i *Item) initializeVertexArrayObject(points []float32) {
	// generate 1 buffer object
	gl.GenBuffers(1, &i.vertexBufferObject)
	// vbo will contain an array of vertices
	gl.BindBuffer(gl.ARRAY_BUFFER, i.vertexBufferObject)
	// fill the buffer with data
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	log.Info("item::initializeVertexArrayObject", i.Name, "position", i.GetPosition(), "rotation", i.GetRotation())

	gl.GenVertexArrays(1, &i.vertexArrayObject)
	gl.BindVertexArray(i.vertexArrayObject)
	// vertex position is treated as vertex attribute index 0 in the fixed function pipeline
	gl.EnableVertexAttribArray(0)
	// tell the pipeline how to interpret the data inside the buffer
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
}
