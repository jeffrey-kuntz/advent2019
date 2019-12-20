package problem3

import (
	"testing"

	"github.com/jeffrey-kuntz/advent2019/math"

	"github.com/stretchr/testify/assert"
)

func TestWire1(t *testing.T) {
	origin := math.NewPoint(0, 0)

	firstWirePoints, _ := GeneratePoints(origin, "R75,D30,R83,U83,L12,D49,R71,U7,L72")
	secondWirePoints, _ := GeneratePoints(origin, "U62,R66,U55,R34,D71,R55,D58,R83")

	intersections := Intersection(firstWirePoints, secondWirePoints)
	distances := Distances(origin, intersections)

	assert.Equal(t, 159, distances[0])
}

func TestWire2(t *testing.T) {
	origin := math.NewPoint(0, 0)

	firstWirePoints, _ := GeneratePoints(origin, "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51")
	secondWirePoints, _ := GeneratePoints(origin, "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7")

	intersections := Intersection(firstWirePoints, secondWirePoints)
	distances := Distances(origin, intersections)

	assert.Equal(t, 135, distances[0])
}
