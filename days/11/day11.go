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
var day = "11"

type Operation func(old int) int
type Test func(val int) int
type Operator func(first int, second int) int

var add Operator = func(first int, second int) int {
	return first + second
}
var multiply Operator = func(first int, second int) int {
	return first * second
}

func reduceWorry(level int) int {
	return level % (17 * 19 * 7 * 11 * 13 * 3 * 5 * 2) // level / 3 for part 1
}

type Monkey struct {
	inspectionCount int
	items           chan int
	operation       Operation
	test            Test
}

var monkeys = make([]*Monkey, 0)

func main() {
	// Construct monkeys
	var index int
	for i := 0; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], "Monkey") {
			index, _ = strconv.Atoi(strings.TrimSuffix(strings.Split(lines[i], " ")[1], ":"))
			monkeys = append(monkeys, &Monkey{items: make(chan int, 50)})
		} else if strings.HasPrefix(lines[i], "  Starting items:") {
			split := strings.Split(strings.TrimPrefix(lines[i], "  Starting items: "), ", ")
			for _, num := range split {
				val, _ := strconv.Atoi(num)
				monkeys[index].items <- val
			}
		} else if strings.HasPrefix(lines[i], "  Operation:") {
			split := strings.Split(strings.TrimPrefix(lines[i], "  Operation: new = "), " ")
			var operator Operator
			if split[1] == "*" {
				operator = multiply
			} else {
				operator = add
			}
			monkeys[index].operation = makeOperationClosure(split[0], split[2], operator)
		} else if strings.HasPrefix(lines[i], "  Test:") {
			divisor, _ := strconv.Atoi(strings.TrimPrefix(lines[i], "  Test: divisible by "))
			trueMonkey, _ := strconv.Atoi(strings.TrimPrefix(lines[i+1], "    If true: throw to monkey "))
			falseMonkey, _ := strconv.Atoi(strings.TrimPrefix(lines[i+2], "    If false: throw to monkey "))
			monkeys[index].test = makeTestClosure(divisor, trueMonkey, falseMonkey)
		}
	}

	const rounds = 10000 // 20 for part 1

	for i := 0; i < rounds; i++ {
		for _, monkey := range monkeys {
			runInspection(monkey)
		}
	}

	inspections := make([]int, len(monkeys))

	for index, monkey := range monkeys {
		inspections[index] = monkey.inspectionCount
	}

	sort.Sort(sort.Reverse(sort.IntSlice(inspections)))
	fmt.Printf("Answer: %v", inspections[0]*inspections[1]) // Part1: 108240, Part2:
}

func runInspection(monkey *Monkey) {
	for len(monkey.items) > 0 {
		newWorryScore := reduceWorry(monkey.operation(<-monkey.items))
		monkeyToThrowTo := monkey.test(newWorryScore)
		monkeys[monkeyToThrowTo].items <- newWorryScore
		monkey.inspectionCount++
	}
}

func makeTestClosure(divisor int, trueMonkey int, falseMonkey int) Test {
	return func(val int) int {
		if val%divisor == 0 {
			return trueMonkey
		} else {
			return falseMonkey
		}
	}
}

func makeOperationClosure(first string, second string, operator Operator) Operation {
	if first == "old" && second == "old" {
		return func(old int) int {
			return operator(old, old)
		}
	} else {
		// all available operations are commutative
		return func(old int) int {
			val, _ := strconv.Atoi(second)
			return operator(old, val)
		}
	}
	panic("We should never get here!")
}

func init() {
	log.SetFlags(0)

	var error error
	lines, error = util.ReadLines(day)

	if error != nil {
		log.Fatal(error)
	}
}
