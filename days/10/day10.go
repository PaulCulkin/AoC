package main

import (
	"fmt"
	"learning/util"
	"log"
	"math"
	"strconv"
	"strings"
)

var lines []string
var day = "10"

type CPU struct {
	signalStrength map[int]int
	cycles         int
	register       int
}

func (cpu *CPU) calculateSignalStrength(cycleCount int) {
	for i := 0; i < cycleCount; i++ {
		cpu.cycles++
		cpu.signalStrength[cpu.cycles] = cpu.register * cpu.cycles
	}
}

func (cpu *CPU) add(val int) {
	cpu.calculateSignalStrength(2)
	cpu.register += val
}

func (cpu *CPU) noop() {
	cpu.calculateSignalStrength(1)
	// no operation
}

func main() {
	cpu := CPU{make(map[int]int), 0, 1}

	for _, instruction := range lines {
		split := strings.Split(instruction, " ")
		switch split[0] {
		case "noop":
			cpu.noop()
		case "addx":
			val, _ := strconv.Atoi(split[1])
			cpu.add(val)
		}
	}

	const (
		start    = 20
		interval = 40
		count    = 6
	)

	total := 0

	for i := 0; i < count; i++ {
		total += cpu.signalStrength[start+(interval*i)]
	}

	fmt.Printf("Part 1: %v\n", total) // 11820

	for i := 0; i < count; i++ {
		for j := 0; j < interval; j++ {
			cycleNum := (interval * i) + j
			register := cpu.signalStrength[cycleNum+1] / (cycleNum + 1)
			if math.Abs(float64(register-j)) <= 1 {
				print("#")
			} else {
				print(" ")
			}
		}
		println()
	}
	// EPJBRKAH
}

func init() {
	log.SetFlags(0)

	var error error
	lines, error = util.ReadLines(day)

	if error != nil {
		log.Fatal(error)
	}
}
