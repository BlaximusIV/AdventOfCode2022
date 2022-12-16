/*
	Part 1 solution was overridden by part 2 solution.
	It was more in line with just pushing and popping from a stack.
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// To ease input parsing, I'm just manually adding the setup here and scrubbing the input
var crates = [][]string{
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
		items := multiPop(movements[1], movements[0])
		crates[movements[2]] = append(crates[movements[2]], items...)
	}
}

func multiPop(crateColumnIndex int, itemCount int) (items []string) {
	stackLength := len(crates[crateColumnIndex])
	firstItemIndex := stackLength - itemCount

	items = crates[crateColumnIndex][firstItemIndex:stackLength]
	crates[crateColumnIndex] = crates[crateColumnIndex][:firstItemIndex]
	return
}

func printFinalCrates() {
	topCrates := ""
	for _, column := range crates {
		topCrates = topCrates + column[len(column)-1]
	}

	fmt.Printf("%v", topCrates)
}
