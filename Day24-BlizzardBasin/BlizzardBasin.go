package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()

	input, _ := os.ReadFile("Input.txt")
	blizzardMap := parseBlizzardMap(string(input))

	// Part 1
	start := Coordinate{0, 1}
	height := len(blizzardMap.Coordinates) - 1
	width := len(blizzardMap.Coordinates[height]) - 1
	goal := Coordinate{height, width - 1}
	routeMinutes := findQuickestRouteTime(blizzardMap, start, goal, 0)
	fmt.Printf("Quickest route in minutes: %d\n", routeMinutes)

	elapsed := time.Since(startTime)
	fmt.Printf("Elapsed time: %v\n", elapsed)
}

func parseBlizzardMap(input string) Map {
	rawMap := [][]string{}

	for _, line := range strings.Split(input, "\n") {
		mapLine := []string{}
		for _, char := range strings.Split(line, "") {
			mapLine = append(mapLine, char)
		}
		rawMap = append(rawMap, mapLine)
	}

	height := len(rawMap) - 2
	width := len(rawMap[0]) - 2
	northBliz, southBliz := [][]bool{}, [][]bool{}
	eastBliz, westBliz := [][]bool{}, [][]bool{}

	for i := 0; i < width; i++ {
		northBliz = append(northBliz, make([]bool, height))
		southBliz = append(southBliz, make([]bool, height))
	}

	for i := 0; i < height; i++ {
		eastBliz = append(eastBliz, make([]bool, width))
		westBliz = append(westBliz, make([]bool, width))
	}

	for i := 1; i <= height; i++ {
		for j := 1; j <= width; j++ {
			char := rawMap[i][j]
			if char == "^" {
				northBliz[j-1][i-1] = true
			} else if char == "v" {
				southBliz[j-1][i-1] = true
			} else if char == "<" {
				westBliz[i-1][j-1] = true
			} else if char == ">" {
				eastBliz[i-1][j-1] = true
			}
		}
	}

	return Map{rawMap, northBliz, southBliz, eastBliz, westBliz}
}

func findQuickestRouteTime(m Map, start Coordinate, goal Coordinate, startTick int) int {
	cycleLength := len(m.East) * len(m.North)
	startState := State{start, startTick}
	startPriority := PriorityItem{startTick + start.EuclidianDistance(goal), startState}

	frontier := PriorityQueue{}
	frontier.Enqueue(startPriority)
	costSoFar := map[State]int{startState: startTick}

	finalState := State{}
	for len(frontier.Vals) > 0 {
		current := frontier.Dequeue()
		if current.Coordinate == goal {
			finalState = current
			break
		}

		moves := m.GetMoves(current.Coordinate, current.Time, start, goal)
		for _, c := range moves {

			currentCostState := State{c.Coordinate, c.Time % cycleLength}
			_, exists := costSoFar[currentCostState]
			if !exists || c.Time < costSoFar[currentCostState] {
				distance := c.Coordinate.EuclidianDistance(goal)
				costSoFar[currentCostState] = c.Time
				frontier.Enqueue(PriorityItem{c.Time + distance, c})
			}
		}
	}

	return int(finalState.Time)
}
