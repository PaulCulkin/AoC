package main

import (
	"fmt"
	"learning/util"
	"log"
	"strconv"
	"strings"
)

var lines []string
var day = "4"

func main() {
	var overlapCount int

	for _, pair := range lines {
		range1, range2 := getRanges(pair)
		if isContained(range1, range2) || isContained(range2, range1) {
			overlapCount++
		}
	}

	fmt.Printf("Part 1: %v\n", overlapCount)

	overlapCount = 0

	for _, pair := range lines {
		range1, range2 := getRanges(pair)
		if isOverlapped(range1, range2) || isOverlapped(range2, range1) {
			overlapCount++
		}
	}

	fmt.Printf("Part 2: %v\n", overlapCount)

}

func isContained(range1 []int, range2 []int) bool {
	return range1[0] <= range2[0] && range1[len(range1)-1] >= range2[len(range2)-1]
}

func isOverlapped(range1 []int, range2 []int) bool {
	return (range2[0] >= range1[0] && range2[0] <= range1[len(range1)-1]) || (range2[len(range2)-1] >= range1[0] && range2[len(range2)-1] <= range1[len(range1)-1])
}

func getRanges(pair string) ([]int, []int) {
	split := strings.Split(pair, ",")
	return getRange(split[0]), getRange(split[1])
}

func getRange(shorthand string) []int {
	split := strings.Split(shorthand, "-")
	start, _ := strconv.Atoi(split[0])
	end, _ := strconv.Atoi(split[1])

	a := make([]int, end-start+1)
	for i := range a {
		a[i] = start + i
	}
	return a
}

func init() {
	log.SetFlags(0)

	var error error
	lines, error = util.ReadLines(day)

	if error != nil {
		log.Fatal(error)
	}
}
