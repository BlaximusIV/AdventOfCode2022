package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// To ease input parsing, I'm just manually adding the setup here, and scrubbing the input
var crates = [][]string{
	// Test Setup
	// {"Z", "N"},
	// {"M", "C", "D"},
	// {"P"},

	// Actual Setup
	{"B", "Q", "C"},
	{"R", "Q", "W", "Z"},
	{"B", "M", "R", "L", "V"},
	{"C", "Z", "H", "V", "T", "W"},
	{"D", "Z", "H", "B", "N", "V", "G"},
	{"H", "N", "P", "C", "J", "F", "V", "Q"},
	{"D", "G", "T", "R", "W", "Z", "S"},
	{"C", "G", "M", "N", "B", "W", "Z", "P"},
	{"N", "J", "B", "M", "W", "Q", "F", "P"},
}

var instructions = [][]int{}

func main() {
	content, err := ioutil.ReadFile("PuzzleInput.txt")

	if err != nil {
		log.Fatal(err)
	}

	parseInstructions(string(content))
	moveCrates()
	printFinalCrates()
}

func parseInstructions(content string) {
	for _, instructionLine := range strings.Split(content, "\r\n") {
		instructionDetails := make([]int, 3)
		instruction := strings.Split(instructionLine, " ")

		crateCount, _ := strconv.Atoi(instruction[1])
		from, _ := strconv.Atoi(instruction[3])
		to, _ := strconv.Atoi(instruction[5])

		// subtract 1 for ease of use with zero-based array index
		instructionDetails[0] = crateCount
		instructionDetails[1] = from - 1
		instructionDetails[2] = to - 1

		instructions = append(instructions, instructionDetails)
	}
}

func moveCrates() {
	for _, movements := range instructions {
		// 0 = count, 1 = from, 2 = to
		for i := 0; i < movements[0]; i++ {
			push(pop(movements[1]), movements[2])
		}
	}
}

// Stack-like methods, because Go doesn't have a stack type
func push(item string, crateColumnIndex int) {
	crates[crateColumnIndex] = append(crates[crateColumnIndex], item)
}

func pop(crateColumnIndex int) (item string) {
	itemIndex := len(crates[crateColumnIndex]) - 1
	item = crates[crateColumnIndex][itemIndex]
	crates[crateColumnIndex] = crates[crateColumnIndex][:itemIndex]
	return
}

func printFinalCrates() {
	for _, column := range crates {
		log.Printf(column[len(column)-1])
	}
}
