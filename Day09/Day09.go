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
	contentString := string(content)

	// Part 1
	tailPositionCount := getTailPositionCount(contentString, 2)
	log.Printf("Tail in %d unique positions with two knots", tailPositionCount)

	// Part 2
	tenKnotTailPositionCount := getTailPositionCount(contentString, 10)
	log.Printf("Tail in %d unique positions with 10 knots", tenKnotTailPositionCount)

	elapsed := time.Since(start)
	log.Printf("Elapsed Time: %s\n", elapsed)
}

func getTailPositionCount(instructionString string, knotCount int) int {
	knots := []coordinate{}
	for i := 0; i < knotCount; i++ {
		knots = append(knots, coordinate{})
	}

	tailVisitedCoordinates := map[coordinate]int{{0, 0}: 1}
	for _, instruction := range strings.Split(instructionString, "\r\n") {
		pieces := strings.Split(instruction, " ")
		direction := pieces[0]
		magnitude, _ := strconv.Atoi(pieces[1])

		for i := 0; i < magnitude; i++ {
			move(&knots[0], direction)
			updateKnotPositions(knots, tailVisitedCoordinates)
		}
	}

	return len(tailVisitedCoordinates)
}

func move(current *coordinate, direction string) {
	switch direction {
	case "R":
		current.X++
	case "U":
		current.Y++
	case "D":
		current.Y--
	case "L":
		current.X--
	}
}

func updateKnotPositions(knots []coordinate, tailVisitedCoordinates map[coordinate]int) {
	for j := range knots[1:] {
		knots[j+1] = getNewPosition(knots[j], knots[j+1])
		if j+2 >= len(knots) {
			tailVisitedCoordinates[knots[j+1]]++
		}
	}
}

// There has to be a better way to model the behavior. I don't know what it is though
func getNewPosition(leader coordinate, follower coordinate) (newCoord coordinate) {
	xDiff := difference(leader.X, follower.X)
	yDiff := difference(leader.Y, follower.Y)
	newCoord = follower

	if xDiff > 1 {
		if leader.X < follower.X {
			newCoord.X--
		} else {
			newCoord.X++
		}

		if yDiff > 1 {
			if leader.Y < follower.Y {
				newCoord.Y--
			} else {
				newCoord.Y++
			}
		} else if yDiff > 0 {
			newCoord.Y = leader.Y
		}
	} else if yDiff > 1 {
		if leader.Y < follower.Y {
			newCoord.Y--
		} else {
			newCoord.Y++
		}

		if xDiff > 1 {
			if leader.X < follower.X {
				newCoord.X--
			} else {
				newCoord.X++
			}
		} else if xDiff > 0 {
			newCoord.X = leader.X
		}
	}

	return
}

func difference(a int, b int) int {
	return int(math.Abs(float64(a - b)))
}
