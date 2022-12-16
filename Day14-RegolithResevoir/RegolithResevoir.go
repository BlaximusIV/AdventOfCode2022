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

	// Part 1
	rockMap, _ := populateMap(string(content))
	placedSandBlocks := simulateSandFall(rockMap)
	log.Printf("Placed sand blocks: %d\n", placedSandBlocks)

	// Part 2
	rockMap, ledgeDepth := populateMap(string(content))
	drawFloor(&rockMap, ledgeDepth)
	placedSandBlocks = simulateSandFall(rockMap)

	log.Printf("Placed sand blocks with floor: %d\n", placedSandBlocks)

	elapsed := time.Since(startTime)
	log.Printf("Elapsed Time: %s\n", elapsed)
}

func populateMap(scan string) (rockMap [200][1200]string, ledgeDepth int) {
	ledges := strings.Split(scan, "\r\n")
	for _, ledge := range ledges {
		directions := strings.Split(ledge, " -> ")
		for i := 0; i < len(directions)-1; i++ {
			a, b := getPoint(directions[i]), getPoint(directions[i+1])
			drawEdge(&rockMap, a, b)
			ledgeDepth = getMax(a.Y, b.Y, ledgeDepth)
		}
	}

	return
}

func drawEdge(rockMap *[200][1200]string, a Point, b Point) {
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

func getMax(a, b, c int) int {
	max := a

	if b > max {
		max = b
	}

	if c > max {
		max = c
	}

	return max
}

func drawFloor(rockMap *[200][1200]string, ledgeDepth int) {
	floorDepth := ledgeDepth + 2
	for i := range rockMap[floorDepth] {
		rockMap[floorDepth][i] = "#"
	}
}

func simulateSandFall(rockMap [200][1200]string) (sandBlocks int) {
	// place block
	spawnPoint := Point{0, 500}
	simulate := true
	for simulate {
		simulate = placeBlock(&rockMap, spawnPoint)

		if simulate {
			sandBlocks++
		}
	}

	return
}

func placeBlock(rockMap *[200][1200]string, location Point) bool {
	// The pile has reached the ceiling hole
	if rockMap[location.Y][location.X] == "O" {
		return false
	}

	sandCanMove := true
	for sandCanMove {
		// It's running into the abyss
		if location.Y+1 >= len(rockMap) {
			return false
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

		rockMap[location.Y][location.X] = "O"
		sandCanMove = false
	}

	return true
}
