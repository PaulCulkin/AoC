package main

import (
	"learning/util"
	"log"
	"strconv"
	"strings"
)

var lines []string
var day = "5"

func main() {

	stacks := [9][]string{}
	var moves [][]int

	for _, row := range lines {
		// Skip empty rows and number column
		if len(row) > 0 && !strings.HasPrefix(row, " 1") {
			if !strings.HasPrefix(row, "move") {
				cleaned := row[1 : len(row)-1]

				for num := range make([]int, 9) {
					elems := cleaned[num*4 : (num*4)+1]
					if elems != " " {
						stacks[num] = append(stacks[num], elems)
					}
				}
			} else if strings.HasPrefix(row, "move") {
				split := strings.Split(row, " ")
				move := make([]int, 3)
				move[0], _ = strconv.Atoi(split[1])
				move[1], _ = strconv.Atoi(split[3])
				move[2], _ = strconv.Atoi(split[5])

				moves = append(moves, move)
			}
		}
	}

	for _, stack := range stacks {
		reverse(stack)
	}

	for _, move := range moves {
		// I feel like this would be a good time to have learnt how to do pointers in Go :(
		fromStack := stacks[move[1]-1]
		toStack := stacks[move[2]-1]
		count := move[0]

		payload := fromStack[len(fromStack)-(count):]

		// Don't reverse for part 2
		//reverse(payload)

		stacks[move[1]-1] = fromStack[:len(fromStack)-count]
		stacks[move[2]-1] = append(toStack, payload...)

	}

	for _, stack := range stacks {
		if len(stack) > 0 {
			print(stack[len(stack)-1])
		}
	}
}

func reverse(stack []string) {
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}
}

func init() {
	log.SetFlags(0)

	var error error
	lines, error = util.ReadLines(day)

	if error != nil {
		log.Fatal(error)
	}
}
