package main

import (
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	elevationValues := getElevationValues()
	elevationValues["S"] = 1
	elevationValues["E"] = 26

	start := time.Now()

	content, _ := os.ReadFile("PuzzleInput.txt")

	regionMap, startLocation, goal := getMapInfo(string(content))
	grid := Grid{len(regionMap), len(regionMap[0]), regionMap, elevationValues}

	shortestPathLength := getShortestPathLength(grid, startLocation, goal)

	// Part 1
	log.Printf("Length of the shortest path to target from current location: %d\n", shortestPathLength)

	elapsed := time.Since(start)
	log.Printf("Elapsed Time: %s\n", elapsed)
}

func getElevationValues() (elevationValues map[string]int) {
	elevationValues = map[string]int{}
	mapValues := "abcdefghijklmnopqrstuvwxyz"
	for i, val := range strings.Split(mapValues, "") {
		elevationValues[val] = i + 1
	}
	return
}

func getMapInfo(input string) (regionMap [][]string, start Coordinate, goal Coordinate) {
	for i, row := range strings.Split(input, "\r\n") {
		rowCharacters := []string{}
		for j, character := range strings.Split(row, "") {
			if character == "S" {
				start = Coordinate{i, j}
			} else if character == "E" {
				goal = Coordinate{i, j}
			}
			rowCharacters = append(rowCharacters, character)
		}
		regionMap = append(regionMap, rowCharacters)
	}

	return
}

func getShortestPathLength(grid Grid, start Coordinate, goal Coordinate) (shortestPathLength int) {
	frontier := PriorityQueue{}
	frontier.Enqueue(PriorityItem{0, start})
	cameFrom := map[Coordinate]Coordinate{}
	costSoFar := map[Coordinate]int{start: 0}

	for len(frontier.Vals) > 0 {
		current := frontier.Dequeue()
		if current == goal {
			break
		}

		for _, c := range grid.neighbors(current) {
			newCost := costSoFar[current] + 1
			_, exists := costSoFar[c]
			if !exists || newCost < costSoFar[c] {
				costSoFar[c] = newCost
				frontier.Enqueue(PriorityItem{newCost, c})
				cameFrom[c] = current
			}
		}
	}

	path := getRoute(cameFrom, start, goal)

	return len(path)
}

func getRoute(vistedLocations map[Coordinate]Coordinate, start Coordinate, goal Coordinate) (path []Coordinate) {
	current := goal

	for current != start {
		path = append(path, current)
		current = vistedLocations[current]
	}

	return
}
