package math

import (
	"fmt"
	"math"
)

// Point3 defines a three-dimensional vertex
type Point3 struct {
	X float64
	Y float64
	Z float64
}

// Magnitude returns the magnitude of the vector
func (p *Point3) Magnitude() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
}

// Normalize returns the normalized vector
func (p *Point3) Normalize() Point3 {
	magnitude := p.Magnitude()
	return Point3{X: p.X / magnitude, Y: p.Y / magnitude, Z: p.Z / magnitude}
}

// Equals returns if two Point3 are the same
func (p *Point3) Equals(q Point3) bool {
	return p.X == q.X && p.Y == q.Y && p.Z == q.Z
}

// ToString shows Point3 as a text representation
func (p *Point3) String() string {
	return fmt.Sprintf("(%f, %f, %f)", p.X, p.Y, p.Z)
}

// Point2 defines a two-dimensional vertex
type Point2 struct {
	X float64
	Y float64
}

// Magnitude returns the magnitude of the vector
func (p *Point2) Magnitude() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

// Normalize returns the normalized vector
func (p *Point2) Normalize() Point2 {
	magnitude := p.Magnitude()
	return Point2{X: p.X / magnitude, Y: p.Y / magnitude}
}

// Equals returns if two Point2 are the same
func (p *Point2) Equals(q Point2) bool {
	return p.X == q.X && p.Y == q.Y
}

// ToString shows Point3 as a text representation
func (p *Point2) String() string {
	return fmt.Sprintf("(%f, %f)", p.X, p.Y)
}
