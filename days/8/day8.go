package main

import (
	"fmt"
	"learning/util"
	"log"
	"strconv"
	"strings"
)

var lines []string
var day = "8"

func main() {

	var grid [][]int

	for _, line := range lines {
		grid = append(grid, parseColumn(line))
	}

	visibleCount := 0
	forAllButTheEdges(grid, countVisible(&visibleCount))

	fmt.Printf("Part 1: %v\n", visibleCount+calculatePerimeter(grid)) // 1851

	currentScore := 0
	forAllButTheEdges(grid, checkHigherScore(&currentScore))

	fmt.Printf("Part 2: %v\n", currentScore) // 574080
}

func parseColumn(line string) []int {
	chars := strings.Split(line, "")
	row := make([]int, len(chars))

	for count, char := range chars {
		intVal, _ := strconv.Atoi(char)
		row[count] = intVal
	}

	return row
}

func countVisible(counter *int) (function func(Tree, [][]int)) {
	return func(tree Tree, grid [][]int) {
		if isVisible(tree, grid) {
			*counter++
		}
	}
}

func checkHigherScore(counter *int) (function func(Tree, [][]int)) {
	return func(tree Tree, grid [][]int) {
		candidateScore := scenicScore(tree, grid)
		if candidateScore > *counter {
			*counter = candidateScore
		}
	}
}

func calculatePerimeter(grid [][]int) int {
	return 2*(len(grid)) + 2*(len(grid[0])-2)
}

func forAllButTheEdges(grid [][]int, function func(Tree, [][]int)) {
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			function(Tree{i, j}, grid)
		}
	}
}

func scenicScore(tree Tree, grid [][]int) int {
	return view(tree, grid, North{}) * view(tree, grid, South{}) * view(tree, grid, East{}) * view(tree, grid, West{})
}

func view[D Direction](tree Tree, grid [][]int, direction D) int {
	height := grid[tree.row][tree.column]
	score := 0

	for i := direction.modifyStart(tree); direction.isEndCondition(i, grid); direction.afterLook(&i) {
		score++
		if direction.selectTree(grid, tree, i) >= height {
			return score
		}
	}

	return score
}

func isVisible(tree Tree, grid [][]int) bool {
	return visible(tree, grid, North{}) || visible(tree, grid, South{}) || visible(tree, grid, East{}) || visible(tree, grid, West{})
}

func visible(tree Tree, grid [][]int, direction Direction) bool {
	height := grid[tree.row][tree.column]

	for i := direction.modifyStart(tree); direction.isEndCondition(i, grid); direction.afterLook(&i) {
		if direction.selectTree(grid, tree, i) >= height {
			return false
		}
	}

	return true
}

type Tree struct {
	row    int
	column int
}

type Direction interface {
	modifyStart(Tree) int
	isEndCondition(int, [][]int) bool
	afterLook(*int)
	selectTree([][]int, Tree, int) int
}

type North struct{}

func (n North) modifyStart(tree Tree) int {
	return tree.row - 1
}
func (n North) isEndCondition(row int, grid [][]int) bool {
	return row >= 0
}
func (n North) afterLook(row *int) {
	*row--
}
func (n North) selectTree(grid [][]int, tree Tree, delta int) int {
	return grid[delta][tree.column]
}

type South struct{}

func (n South) modifyStart(tree Tree) int {
	return tree.row + 1
}
func (n South) isEndCondition(row int, grid [][]int) bool {
	return row < len(grid)
}
func (n South) afterLook(row *int) {
	*row++
}
func (n South) selectTree(grid [][]int, tree Tree, delta int) int {
	return grid[delta][tree.column]
}

type East struct{}

func (n East) modifyStart(tree Tree) int {
	return tree.column + 1
}
func (n East) isEndCondition(column int, grid [][]int) bool {
	return column < len(grid[0])
}
func (n East) afterLook(row *int) {
	*row++
}
func (n East) selectTree(grid [][]int, tree Tree, delta int) int {
	return grid[tree.row][delta]
}

type West struct{}

func (n West) modifyStart(tree Tree) int {
	return tree.column - 1
}
func (n West) isEndCondition(column int, grid [][]int) bool {
	return column >= 0
}
func (n West) afterLook(row *int) {
	*row--
}
func (n West) selectTree(grid [][]int, tree Tree, delta int) int {
	return grid[tree.row][delta]
}

func init() {
	log.SetFlags(0)

	var error error
	lines, error = util.ReadLines(day)

	if error != nil {
		log.Fatal(error)
	}
}
