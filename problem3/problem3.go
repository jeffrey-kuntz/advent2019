package problem3

import (
	"bufio"
	"errors"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/jeffrey-kuntz/advent2019/math"
)

// LoadWires loads the wire data from a file
func LoadWires(filename string) (string, string, error) {
	infile, err := os.Open(filename)
	if err != nil {
		return "", "", err
	}

	scanner := bufio.NewScanner(infile)
	scanner.Scan()
	firstWire := scanner.Text()
	scanner.Scan()
	secondWire := scanner.Text()

	return firstWire, secondWire, nil
}

// GeneratePoints generates a set of points gives a comma seperated string of vectors
func GeneratePoints(origin *math.Point, directionLine string) ([]*math.Point, error) {
	vectors := strings.Split(directionLine, ",")

	var pointList []*math.Point

	// Start the location at the origin
	currentLocation := math.NewPoint(origin.X, origin.Y)

	for _, vector := range vectors {
		direction := string(vector[0])
		distance, err := strconv.Atoi(vector[1:])
		if err != nil {
			return nil, errors.New("vector is not parsable: " + vector)
		}

		switch direction {
		case "U":
			pointList = addYPoints(pointList, currentLocation.X, currentLocation.Y, currentLocation.Y+distance)
			currentLocation.Y = currentLocation.Y + distance
		case "D":
			pointList = addYPoints(pointList, currentLocation.X, currentLocation.Y-distance, currentLocation.Y)
			currentLocation.Y = currentLocation.Y - distance
		case "L":
			pointList = addXPoints(pointList, currentLocation.Y, currentLocation.X-distance, currentLocation.X)
			currentLocation.X = currentLocation.X - distance
		case "R":
			pointList = addXPoints(pointList, currentLocation.Y, currentLocation.X, currentLocation.X+distance)
			currentLocation.X = currentLocation.X + distance
		default:
			return nil, errors.New("vector is not parsable: " + vector)
		}
	}

	return pointList, nil
}

func addYPoints(pointList []*math.Point, currentX int, startIndex int, endIndex int) []*math.Point {
	for i := startIndex + 1; i <= endIndex; i++ {
		pointList = append(pointList, math.NewPoint(currentX, i))
	}

	return pointList
}

func addXPoints(pointList []*math.Point, currentY int, startIndex int, endIndex int) []*math.Point {
	for i := startIndex + 1; i <= endIndex; i++ {
		pointList = append(pointList, math.NewPoint(i, currentY))
	}

	return pointList
}

// Intersection finds the intersection of points between the first and second sets
func Intersection(firstPoints []*math.Point, secondPoints []*math.Point) []*math.Point {
	var points []*math.Point

	for _, firstPoint := range firstPoints {
		for _, secondPoint := range secondPoints {
			isEqual := firstPoint.IsEqual(secondPoint)

			if isEqual {
				points = append(points, firstPoint)
			}
		}
	}

	return points
}

// Distances gets a sorted list of distances from the origin points
func Distances(origin *math.Point, points []*math.Point) []int {
	var distances []int
	for _, point := range points {
		distance := origin.GetManhattenDistance(point)
		distances = append(distances, distance)
	}

	sort.Ints(distances)

	return distances
}
