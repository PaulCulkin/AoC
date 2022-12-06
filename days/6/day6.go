package main

import (
	"fmt"
	"learning/util"
	"log"
)

var line []string
var day = "6"

func main() {
	packetLength := 4 // change to 14 for part 2
	packet := line[:packetLength]

	for count, char := range line {
		if isUnique(packet) {
			fmt.Printf("Answer: %v\n", count)
			return
		} else {
			packet[count%packetLength] = char
		}
	}
}

func isUnique(packet []string) bool {
	for i := 0; i < len(packet); i++ {
		for j := i + 1; j < len(packet); j++ {
			if packet[i] == packet[j] {
				return false
			}
		}
	}
	return true
}

func init() {
	log.SetFlags(0)

	var error error
	line, error = util.ReadLine(day)

	if error != nil {
		log.Fatal(error)
	}
}
