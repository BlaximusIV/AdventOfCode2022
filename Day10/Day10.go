package main

import (
	"log"
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

	checkedCycles := []int{20, 60, 100, 140, 180, 220}
	signalStrengths := executeInstructions(instructions, checkedCycles)

	// Part 1
	log.Printf("Strengths: %v", signalStrengths)
	sum := 0
	for _, val := range signalStrengths {
		sum += val
	}
	log.Printf("Sum of strengths: %d\n", sum)

	elapsed := time.Since(start)
	log.Printf("Elapsed Time: %s\n", elapsed)
}

func executeInstructions(instructions []string, checkedCycles []int) (signalStrengths []int) {
	registerX := 1
	cycleCount := 1
	instructionType := ""
	instructionParam := 0
	var instructionTtl int

	for len(instructions) > 0 {
		if contains(checkedCycles, cycleCount) {
			signalStrengths = append(signalStrengths, cycleCount*registerX)
		}
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

func contains(slice []int, target int) (contains bool) {
	for _, item := range slice {
		if item == target {
			contains = true
			return
		}
	}
	return
}
