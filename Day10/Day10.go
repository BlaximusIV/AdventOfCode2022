package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

var instructionLifetimes = map[string]int{"noop": 1, "addx": 2}

func main() {
	start := time.Now()

	content, _ := os.ReadFile("PuzzleInput.txt")
	instructions := strings.Split(string(content), "\r\n")

	signalStrengths, crtScreen := getExecutionResults(instructions)

	// Part 1
	printSummedStrengths(signalStrengths)

	// Part 2
	drawScreen(crtScreen)

	elapsed := time.Since(start)
	log.Printf("Elapsed Time: %s\n", elapsed)
}

// Per puzzle input, crt screen is of fixed size
func getExecutionResults(instructions []string) (signalStrengths []int, crtScreen [6][40]string) {
	registerX := 1
	cycleCount := 1
	instructionType := ""
	instructionParam := 0
	var instructionTtl int
	crtRenderRow := 0
	for len(instructions) > 0 {
		signalStrengths = append(signalStrengths, cycleCount*registerX)

		rowPosition := (cycleCount - 1) % 40
		if cycleCount != 1 && rowPosition == 0 {
			crtRenderRow++
		}

		characterToDraw := "."
		if difference(registerX, rowPosition) <= 1 {
			characterToDraw = "#"
		}

		crtScreen[crtRenderRow][rowPosition] = characterToDraw

		if instructionTtl < 1 {
			instructionParts := strings.Split(instructions[0], " ")
			instructions = instructions[1:]
			instructionType, instructionTtl = instructionParts[0], instructionLifetimes[instructionParts[0]]
			if instructionType == "addx" {
				instructionParam, _ = strconv.Atoi(instructionParts[1])
			}
		}

		cycleCount++
		instructionTtl--

		if instructionType == "addx" && instructionTtl == 0 {
			registerX += instructionParam
		}
	}

	return
}

func difference(a int, b int) int {
	return int(math.Abs(float64(a - b)))
}

func printSummedStrengths(strengths []int) {
	// Specific indexes to check for puzzle solution given in puzzle prompt
	checkedStrengths := []int{strengths[19], strengths[59], strengths[99], strengths[139], strengths[179], strengths[219]}
	sum := 0
	for _, val := range checkedStrengths {
		sum += val
	}
	log.Printf("Sum of strengths: %d\n", sum)
}

func drawScreen(crtScreen [6][40]string) {
	for _, row := range crtScreen {
		log.Printf("%v\r\n", row)
	}
}
