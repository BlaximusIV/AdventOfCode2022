package main

import (
	"log"
	"os"
	"strings"
	"time"
)

var elevationValues map[string]int
var impassable = 100_000_000

type Coordinate struct {
	Y, X int
}

func main() {
	elevationValues["S"] = 1
	elevationValues["E"] = 26

	start := time.Now()

	content, _ := os.ReadFile("PuzzleTestInput.txt")

	regionMap, startLocation, goal := populateRegionMap(string(content))

	shortestPathLength := getShortestPathLength(regionMap, startLocation, goal)

	elapsed := time.Since(start)
	log.Printf("Elapsed Time: %s\n", elapsed)
}

func populateElevationValues() {
	mapValues := "abcdefghijklmnopqrstuvwxyz"
	for i, val := range strings.Split(mapValues, "") {
		elevationValues[val] = i + 1
	}
}

func populateRegionMap(input string) (regionMap [][]string, startCoordinate Coordinate, goalCoordinate Coordinate) {
	for i, row := range strings.Split(input, "\r\n") {
		rowCharacters := []string{}
		for j, character := range strings.Split(row, "") {
			if character == "S" {
				startCoordinate = Coordinate{i, j}
			} else if character == "E" {
				goalCoordinate = Coordinate{i, j}
			}
			rowCharacters = append(rowCharacters, character)
		}
		regionMap = append(regionMap, rowCharacters)
	}

	return
}

func getShortestPathLength(regionMap [][]string, startCoordinate Coordinate, goalCoordinate Coordinate)
