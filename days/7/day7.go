package main

import (
	"fmt"
	"learning/util"
	"log"
	"sort"
	"strconv"
	"strings"
)

var lines []string
var day = "7"

type Node interface {
	getSize() int
}

type Dir struct {
	name     string
	children []Node
	parent   *Dir
}

func (dir *Dir) getSize() int {
	sum := 0
	for _, child := range dir.children {
		sum += child.getSize()
	}
	return sum
}

type File struct {
	name   string
	size   int
	parent *Dir
}

func (file *File) getSize() int {
	return file.size
}

func parseLine(line string, dir *Dir) *Dir {
	if strings.HasPrefix(line, "$ ") {
		return parseCommand(strings.TrimPrefix(line, "$ "), dir)
	} else {
		// create files and directories
		split := strings.Split(line, " ")
		if split[0] == "dir" {
			dir.children = append(dir.children, &Dir{name: split[1], parent: dir, children: []Node{}})
		} else {
			size, _ := strconv.Atoi(split[0])
			dir.children = append(dir.children, &File{name: split[1], parent: dir, size: size})
		}
	}
	return dir
}

func parseCommand(command string, dir *Dir) *Dir {
	switch {
	case strings.HasPrefix(command, "cd"):
		return handleDirChange(strings.Split(command, " ")[1], dir)
	case strings.HasPrefix(command, "ls"):
		// listing dir
	}
	return dir
}

func handleDirChange(newDir string, dir *Dir) *Dir {
	switch newDir {
	case "..":
		return dir.parent
	case "/":
		return getRoot(dir)
	default:
		{
			for _, child := range dir.children {
				switch child.(type) {
				case *Dir:
					{
						dir := child.(*Dir)
						if dir.name == newDir {
							return dir
						}
					}
				}
			}
			// implicit create if dir not found? i.e. do we ever cd to a directory before ls it?
			newDir := &Dir{name: newDir, parent: dir, children: []Node{}}
			dir.children = append(dir.children, newDir)
			return newDir
		}
	}
}

func getRoot(currentDir *Dir) *Dir {
	for currentDir.parent != nil {
		currentDir = currentDir.parent
	}
	return currentDir
}

func main() {
	var currentDir = &Dir{
		name:     "root",
		children: []Node{},
	}

	for _, line := range lines {
		currentDir = parseLine(line, currentDir)
	}

	root := getRoot(currentDir)

	fmt.Printf("Part 1: %v\n", calculateSumSizeBelow(100000, root))

	spaceRequired := root.getSize() + 30000000 - 70000000

	var sizes []int

	sizes = populateDirSizes(root, spaceRequired)

	sort.Ints(sizes)

	fmt.Printf("Part 2: %v\n", sizes[0])
}

func populateDirSizes(root *Dir, required int) []int {
	var sizes []int

	if root.getSize() >= required {
		sizes = append(sizes, root.getSize())

		for _, node := range root.children {
			if dir, isDir := node.(*Dir); isDir {
				sizes = append(sizes, populateDirSizes(dir, required)...)
			}
		}
	}

	return sizes
}

func calculateSumSizeBelow(maxSize int, root *Dir) int {
	sum := 0

	if root.getSize() <= maxSize {
		sum += root.getSize()
	}

	for _, node := range root.children {
		if dir, isDir := node.(*Dir); isDir {
			sum += calculateSumSizeBelow(maxSize, dir)
		}
	}

	return sum
}

func init() {
	log.SetFlags(0)

	var error error
	lines, error = util.ReadLines(day)

	if error != nil {
		log.Fatal(error)
	}
}
