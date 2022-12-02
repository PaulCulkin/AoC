package main

import (
	"fmt"
	"learning/util"
	"log"
	"strings"
)

var lines []string
var day = "2"

func main() {
	var totalScore int

	for _, round := range lines {
		split := strings.Split(round, " ")
		totalScore += getScore(split[0], split[1])
	}

	fmt.Printf("Part 1: %v\n", totalScore)

	totalScore = 0

	for _, round := range lines {
		split := strings.Split(round, " ")
		totalScore += getScore(split[0], transform(split[1], split[0]))
	}

	fmt.Printf("Part 2: %v\n", totalScore)
}

func transform(result string, opponent string) string {
	if result == "X" {
		if opponent == "A" {
			return "Z"
		}
		if opponent == "B" {
			return "X"
		}
		if opponent == "C" {
			return "Y"
		}
	}
	if result == "Y" {
		if opponent == "A" {
			return "X"
		}
		if opponent == "B" {
			return "Y"
		}
		if opponent == "C" {
			return "Z"
		}
	}
	if result == "Z" {
		if opponent == "A" {
			return "Y"
		}
		if opponent == "B" {
			return "Z"
		}
		if opponent == "C" {
			return "X"
		}
	}
	return ""
}

func getScore(you string, me string) int {
	var total = 0

	if me == "X" {
		total += 1

		if you == "A" {
			total += 3
		}
		if you == "B" {
			total += 0
		}
		if you == "C" {
			total += 6
		}
	}
	if me == "Y" {
		total += 2

		if you == "A" {
			total += 6
		}
		if you == "B" {
			total += 3
		}
		if you == "C" {
			total += 0
		}
	}
	if me == "Z" {
		total += 3

		if you == "A" {
			total += 0
		}
		if you == "B" {
			total += 6
		}
		if you == "C" {
			total += 3
		}
	}
	return total
}

func init() {
	log.SetFlags(0)

	var error error
	lines, error = util.ReadLines(day)

	if error != nil {
		log.Fatal(error)
	}
}
