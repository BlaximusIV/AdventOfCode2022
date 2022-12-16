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

	// Part 2
	// We're gonna brute-force this cause it'll be easy and I need to move on to other things
	minIdealPathLength := shortestPathLength
	aCoordinates := getACoordinates(grid)
	for _, coord := range aCoordinates {
		pathLength := getShortestPathLength(grid, coord, goal)
		if pathLength == 0 {
			continue
		} else if pathLength < minIdealPathLength {
			minIdealPathLength = pathLength
		}
	}

	log.Printf("Length of the shortest path to target from ideal location: %d\n", minIdealPathLength)

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

		for _, c := range grid.Neighbors(current) {
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
	_, exists := vistedLocations[goal]
	if !exists {
		return
	}

	current := goal

	for current != start {
		path = append(path, current)
		current = vistedLocations[current]
	}

	return
}

func getACoordinates(grid Grid) (coordinates []Coordinate) {
	for i := range grid.RegionMap {
		for j := range grid.RegionMap[i] {
			if grid.RegionMap[i][j] == "a" {
				coordinates = append(coordinates, Coordinate{i, j})
			}
		}
	}

	return
}
