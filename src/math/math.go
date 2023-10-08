package math

import (
	"math"
)

// Point3 defines a three-dimensional vertex
type Point3 struct {
	X float64
	Y float64
	Z float64
}

func (p Point3) magnitude() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)

}

func (p Point3) normalize() Point3 {
	magnitude := p.magnitude()
	return Point3{X: p.X / magnitude, Y: p.Y / magnitude, Z: p.Z / magnitude}
}

func (p Point3) equals(q Point3) bool {
	return p.X == q.X && p.Y == q.Y && p.Z == q.Z
}

// Point2 defines a two-dimensional vertex
type Point2 struct {
	X float64
	Y float64
}

func (p Point2) magnitude() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)

}

func (p Point2) normalize() Point2 {
	magnitude := p.magnitude()
	return Point2{X: p.X / magnitude, Y: p.Y / magnitude}
}

func (p Point2) equals(q Point2) bool {
	return p.X == q.X && p.Y == q.Y
}
