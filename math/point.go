package math

import "math"

// Point defines a spot on a cartesian grid
type Point struct {
	X int
	Y int
}

// NewPoint creates a point
func NewPoint(x int, y int) *Point {
	p := Point{X: x, Y: y}
	return &p
}

// IsEqual checks if our point is equal to passed in point
func (p *Point) IsEqual(op *Point) bool {
	if op == nil {
		return false
	}

	if p.X == op.X && p.Y == op.Y {
		return true
	}

	return false
}

// GetManhattenDistance finds the distance from one point to another
func (p *Point) GetManhattenDistance(op *Point) int {
	return int(math.Abs(float64(op.X)-float64(p.X)) + math.Abs(float64(op.Y)-float64(p.Y)))
}
