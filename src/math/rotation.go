package math

import (
	"fmt"
	"math"
)

// Rotation matrix =
//   | cos(r)  -sin(r) |
//   | sin(r)   cos(r) |

// Inner representation for fast access
//   | a b |
//   | c d |

// Rotation defines a three-dimensional vertex
type Rotation struct {
	Angle float64

	A float64
	B float64
	C float64
	D float64
}

// GetAngle returns the rotation angle
func (r *Rotation) GetAngle() float64 {
	return r.Angle
}

// Rotate rotates the rotation matrix
func (r *Rotation) Rotate(dRot float64) {
	r.Update(r.Angle + dRot)
}

// Update updates the rotation matrix
func (r *Rotation) Update(rot float64) {
	r.Angle = rot
	r.A = math.Cos(rot)
	r.B = -math.Sin(rot)
	r.C = math.Sin(rot)
	r.D = math.Cos(rot)
}

// Add adds an angle to a rotation
func (r *Rotation) Add(rot *Rotation) {
	r.Update(r.Angle + rot.Angle)
}

// TransformPosition transforms a position using the rotation matrix
func (r *Rotation) TransformPosition(p *Point2) *Point2 {
	return &Point2{X: p.X*r.A + p.Y*r.B, Y: p.X*r.C + p.Y*r.D}
}

func (r *Rotation) String() string {
	return fmt.Sprintf("(%f)", r.Angle)
}
