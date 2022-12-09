package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type coordinate struct {
	X, Y int
}

func main() {
	start := time.Now()

	content, _ := os.ReadFile("PuzzleInput.txt")

	tailPositionCount := getTailPositionCount(string(content))

	// Part 1
	log.Printf("Tail in %d unique positions", tailPositionCount)

	elapsed := time.Since(start)
	log.Printf("Elapsed Time: %s\n", elapsed)
}

// This should probably be split into multiple methods and spaced for readability
func getTailPositionCount(instructionString string) int {
	headPosition := coordinate{0, 0}
	tailPosition := coordinate{0, 0}
	tailVisitedCoordinates := map[coordinate]int{{0, 0}: 1}
	for _, instruction := range strings.Split(instructionString, "\r\n") {
		pieces := strings.Split(instruction, " ")
		// new coord, head coord tail coord
		direction := pieces[0]
		magnitude, _ := strconv.Atoi(pieces[1])
		switch direction {
		case "R":
			for i := 0; i < magnitude; i++ {
				newPosition := coordinate{headPosition.X + 1, headPosition.Y}
				tailPosition = moveTail(newPosition, headPosition, tailPosition, tailVisitedCoordinates)
				headPosition = newPosition
			}
		case "U":
			for i := 0; i < magnitude; i++ {
				newPosition := coordinate{headPosition.X, headPosition.Y + 1}
				tailPosition = moveTail(newPosition, headPosition, tailPosition, tailVisitedCoordinates)
				headPosition = newPosition
			}
		case "D":
			for i := 0; i < magnitude; i++ {
				newPosition := coordinate{headPosition.X, headPosition.Y - 1}
				tailPosition = moveTail(newPosition, headPosition, tailPosition, tailVisitedCoordinates)
				headPosition = newPosition
			}
		case "L":
			for i := 0; i < magnitude; i++ {
				newPosition := coordinate{headPosition.X - 1, headPosition.Y}
				tailPosition = moveTail(newPosition, headPosition, tailPosition, tailVisitedCoordinates)
				headPosition = newPosition
			}
		}

	}
	return len(tailVisitedCoordinates)
}

// That's a lot of params. Might be good to split it up somehow
func moveTail(newPosition coordinate, headPosition coordinate, tailPosition coordinate, visitedCoordinates map[coordinate]int) coordinate {
	hasMovedTooFar := math.Abs(float64(newPosition.X-tailPosition.X)) > 1 || math.Abs(float64(newPosition.Y-tailPosition.Y)) > 1
	if hasMovedTooFar {
		visitedCoordinates[headPosition]++
		return headPosition
	}

	return tailPosition
}
