package main

import (
	"fmt"
	"learning/util"
	"log"
	"strconv"
	"strings"
)

var lines []string
var day = "9"

const numKnots = 2 // change to 10 for part 2

type Location struct {
	x int
	y int
}

func (l Location) Equals(location Location) bool {
	return l.x == location.x && l.y == location.y
}

type Rope struct {
	knots [numKnots]Location
}

func main() {

	var locations []Location
	rope := new(Rope)

	for _, movement := range lines {
		direction, count := parseCommand(movement)
		locations = append(locations, applyMovement(direction, count, rope)...)
	}

	fmt.Printf("Part 1: %v\n", countUniqueLocations(locations)) // 6314, 2504
}

func parseCommand(command string) (string, int) {
	split := strings.Split(command, " ")
	count, _ := strconv.Atoi(split[1])
	return split[0], count
}

func applyMovement(direction string, count int, rope *Rope) []Location {
	var tailPositions []Location

	for i := 0; i < count; i++ {
		switch direction {
		case "U":
			rope.knots[0].y++
		case "D":
			rope.knots[0].y--
		case "L":
			rope.knots[0].x--
		case "R":
			rope.knots[0].x++
		}
		for i := 1; i < len(rope.knots); i++ {
			rope.knots[i] = calculateTail(rope.knots[i-1], rope.knots[i])
		}
		tailPositions = append(tailPositions, rope.knots[numKnots-1])
	}

	return tailPositions
}

func calculateTail(head Location, tail Location) Location {
	if touching(head, tail) {
		return tail
	} else {
		if result, plane := isOnSamePlane(head, tail); result {
			switch plane {
			case "X":
				if head.x > tail.x {
					return Location{x: tail.x + 1, y: tail.y}
				} else {
					return Location{x: tail.x - 1, y: tail.y}
				}
			case "Y":
				if head.y > tail.y {
					return Location{x: tail.x, y: tail.y + 1}
				} else {
					return Location{x: tail.x, y: tail.y - 1}
				}
			default:
				panic("We are on the same plane but it's neither X nor Y - such dimension shift like wow!")
			}
		} else {
			// Diagon alley
			newLocation := tail
			if head.x > tail.x {
				newLocation.x++
			} else {
				newLocation.x--
			}
			if head.y > tail.y {
				newLocation.y++
			} else {
				newLocation.y--
			}
			return newLocation
		}
	}
}

func isOnSamePlane(head Location, tail Location) (bool, string) {
	if head.x == tail.x {
		return true, "Y"
	}
	if head.y == tail.y {
		return true, "X"
	}
	return false, ""
}

func touching(tail Location, head Location) bool {
	touchingLocations := []Location{
		{
			x: head.x,
			y: head.y,
		},
		{
			x: head.x + 1,
			y: head.y,
		},
		{
			x: head.x + 1,
			y: head.y - 1,
		},
		{
			x: head.x,
			y: head.y - 1,
		},
		{
			x: head.x - 1,
			y: head.y - 1,
		},
		{
			x: head.x - 1,
			y: head.y,
		},
		{
			x: head.x - 1,
			y: head.y + 1,
		},
		{
			x: head.x,
			y: head.y + 1,
		},
		{
			x: head.x + 1,
			y: head.y + 1,
		},
	}

	for _, location := range touchingLocations {
		if tail.Equals(location) {
			return true
		}
	}

	return false
}

func countUniqueLocations(locations []Location) int {
	uniqueLocations := make(map[string]Location)

	for _, location := range locations {
		uniqueLocations[fmt.Sprintf("%v-%v", location.x, location.y)] = location
	}

	return len(uniqueLocations)
}

func init() {
	log.SetFlags(0)

	var error error
	lines, error = util.ReadLines(day)

	if error != nil {
		log.Fatal(error)
	}
}
