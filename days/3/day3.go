package main

import (
	"fmt"
	"learning/util"
	"log"
	"strings"
)

var lines []string
var day = "3"

func main() {
	var priorityScore int

	for _, bag := range lines {
		priorityScore += getPriority(getCommonItems(getCompartments(bag)))
	}

	fmt.Printf("Part 1: %v\n", priorityScore)

	priorityScore = 0

	groupMembers := [3]string{}

	for count, bag := range lines {
		memberNumber := count % 3
		groupMembers[memberNumber] = bag
		if memberNumber == 2 {
			priorityScore += getPriority(getCommonItems(getCommonItems(groupMembers[0], groupMembers[1]), groupMembers[2]))
		}
	}

	fmt.Printf("Part 2: %v\n", priorityScore)

}

func getCompartments(bag string) (string, string) {
	runes := []rune(bag)
	return string(runes[:len(runes)/2]), string(runes[len(runes)/2:])
}

func getCommonItems(compartment1 string, compartment2 string) string {
	var items []string
	for i := range compartment1 {
		for j := range compartment2 {
			if compartment1[i] == compartment2[j] {
				items = append(items, compartment1[i:i+1])
			}
		}
	}
	return strings.Join(items, "")
}

func getPriority(item string) int {
	i := int(item[0]) - int('a') + 1
	if i <= 0 {
		i = int(item[0]) - int('A') + 1 + 26
	}
	return i
}

func init() {
	log.SetFlags(0)

	var error error
	lines, error = util.ReadLines(day)

	if error != nil {
		log.Fatal(error)
	}
}
