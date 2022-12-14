package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	Y, X int
}

func main() {
	startTime := time.Now()

	content, _ := os.ReadFile("PuzzleInput.txt")

	var rockMap [200][600]string
	populateMap(&rockMap, string(content))

	placedSandBlocks := simulateSandFall(rockMap)

	// Part 1
	log.Printf("Placed sand blocks: %d\n", placedSandBlocks)

	elapsed := time.Since(startTime)
	log.Printf("Elapsed Time: %s\n", elapsed)
}

func populateMap(rockMap *[200][600]string, scan string) {
	ledges := strings.Split(scan, "\r\n")
	for _, ledge := range ledges {
		directions := strings.Split(ledge, " -> ")
		for i := 0; i < len(directions)-1; i++ {
			drawEdge(rockMap, getPoint(directions[i]), getPoint(directions[i+1]))
		}
	}
}

func drawEdge(rockMap *[200][600]string, a Point, b Point) {
	rockMap[a.Y][a.X] = "#"
	for a != b {
		// transform
		if a.Y > b.Y {
			a.Y--
		} else if a.Y < b.Y {
			a.Y++
		} else if a.X < b.X {
			a.X++
		} else {
			a.X--
		}

		rockMap[a.Y][a.X] = "#"
	}
}

func getPoint(point string) Point {
	xy := strings.Split(point, ",")
	x, _ := strconv.Atoi(xy[0])
	y, _ := strconv.Atoi(xy[1])
	return Point{y, x}
}

func simulateSandFall(rockMap [200][600]string) int {
	blockPositions := []Point{}
	// place block
	spawnPoint := Point{0, 500}
	simulate := true
	for simulate {
		cont, location := placeBlock(&rockMap, spawnPoint)
		simulate = cont
		if cont {
			blockPositions = append(blockPositions, location)
		}
	}

	return len(blockPositions)
}

func placeBlock(rockMap *[200][600]string, location Point) (bool, Point) {
	sandCanMove := true
	for sandCanMove {
		// It's running into the abyss
		if location.Y+1 >= len(rockMap) {
			return false, location
		}

		// If it can move down
		if rockMap[location.Y+1][location.X] == "" {
			location.Y++
			continue
		}

		// Can move left
		if rockMap[location.Y+1][location.X-1] == "" {
			location.Y++
			location.X--
			continue
		}

		// Can move right
		if rockMap[location.Y+1][location.X+1] == "" {
			location.Y++
			location.X++
			continue
		}

		// Time to settle a new grain
		rockMap[location.Y][location.X] = "O"
		sandCanMove = false
	}

	return true, location
}
