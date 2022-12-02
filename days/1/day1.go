package main

import (
	"fmt"
	"learning/util"
	"log"
	"sort"
	"strconv"
)

var lines []string
var day = "1"

func main() {
	var calories []int
	var calorieCount int

	for _, calorie := range lines {
		if calorie == "" {
			calories = append(calories, calorieCount)
			calorieCount = 0
		} else {
			intValue, _ := strconv.Atoi(calorie)
			calorieCount += intValue
		}
	}

	sort.Ints(calories)

	fmt.Printf("Part 1: %v\n", calories[len(calories)-1])
	fmt.Printf("Part 2: %v\n", calories[len(calories)-1]+calories[len(calories)-2]+calories[len(calories)-3])
}

func init() {
	log.SetFlags(0)

	var error error
	lines, error = util.ReadLines(day)

	if error != nil {
		log.Fatal(error)
	}
}
