package main

import (
	"log"
	"math"
	"os"
	"time"
)

const ChamberWidth = 7

// In case I want to switch to another delimeter
const SpaceDelimeter = " "

var NewRow = [7]string{SpaceDelimeter, SpaceDelimeter, SpaceDelimeter, SpaceDelimeter, SpaceDelimeter, SpaceDelimeter, SpaceDelimeter}

type Block struct {
	Points []Point
}

type State struct {
	BlockTop                 [50][7]string
	WindIndex, NextIteration int
}

func (b Block) FindLeft() int {
	left := b.Points[0].X
	for i := 1; i < len(b.Points); i++ {
		if b.Points[i].X < left {
			left = b.Points[i].X
		}
	}
	return left
}

func (b Block) FindRight() int {
	right := b.Points[0].X
	for i := 1; i < len(b.Points); i++ {
		if b.Points[i].X > right {
			right = b.Points[i].X
		}
	}
	return right
}

func (b Block) FindBottom() int {
	bottom := b.Points[0].Y
	for i := 1; i < len(b.Points); i++ {
		if b.Points[i].Y > bottom {
			bottom = b.Points[i].Y
		}
	}
	return bottom
}

type Point struct {
	Y, X int
}

func main() {
	startTime := time.Now()

	content, _ := os.ReadFile("Input.txt")

	// Part1
	iterations := 2022
	towerHeight, _ := getSimulatedTowerHeight(string(content), iterations)
	log.Printf("The tower is %d blocks tall \n", towerHeight)

	// Part2
	// One Trillion, find the cycle
	iterations = 3500
	_, states := getSimulatedTowerHeight(string(content), iterations)
	cycleLength, state := findCycle(states)
	heightOne, _ := getSimulatedTowerHeight(string(content), cycleLength)
	heightTwo, _ := getSimulatedTowerHeight(string(content), cycleLength*2)
	cycleHeight := heightTwo - heightOne

	// How many cycles does it take to get to the target?
	const targetCount = 2022 //1_000_000_000_000
	cycleCount := targetCount / cycleLength
	remainder := targetCount - (cycleCount * cycleLength)

	// Find the tower height of the remainder and add it
	h, _ := getSimulatedTowerHeightFromState(string(content), state.WindIndex, targetCount-remainder, targetCount, state.BlockTop[:])
	totalHeight := (h - remainder) + (cycleCount * cycleHeight)
	log.Printf("The tower is %d blocks tall \n", totalHeight)

	elapsed := time.Since(startTime)
	log.Printf("Elapsed Time: %s\n", elapsed)
}

func getSimulatedTowerHeightFromState(directions string, directionsIndex, startIteration, endIteration int, state [][7]string) (height int, states []State) {
	chamber := state
	states = []State{}

	for i := startIteration; i < endIteration; i++ {
		// Spawn block
		block := spawnBlock(i)

		// Make room
		const spawnMargin = 3
		difference := spawnMargin - findHighestBlock(chamber) // How much space we have

		// Make sure we've got the correct number up front.
		if difference < 0 {
			chamber = chamber[difference*-1:]
		} else if difference > 0 {
			insertRows(&chamber, difference)
		}

		rowsToInsert := block.FindBottom() + 1
		insertRows(&chamber, rowsToInsert)

		// loop through block movements until rest
		canLower := true
		for canLower {
			direction := string(directions[directionsIndex])
			tryShift(chamber, &block, direction)
			canLower = tryShift(chamber, &block, "^")

			// Ensure index
			directionsIndex++
			if directionsIndex == len(directions) {
				directionsIndex = 0
			}
		}

		// write to chamber
		for _, val := range block.Points {
			chamber[val.Y][val.X] = "#"
		}

		// Add completed state
		states = append(states, State{getTowerPositionsState(chamber), directionsIndex, i + 1})
	}

	height = len(chamber) - findHighestBlock(chamber)
	return
}

func getSimulatedTowerHeight(directions string, iterations int) (height int, states []State) {
	// the simulated chamber
	chamber := [][7]string{NewRow, NewRow, NewRow}
	states = []State{}
	directionIndex := 0
	for i := 0; i < iterations; i++ {
		// Spawn block
		block := spawnBlock(i)

		// Make room
		const spawnMargin = 3
		difference := spawnMargin - findHighestBlock(chamber) // How much space we have

		// Make sure we've got the correct number up front.
		if difference < 0 {
			chamber = chamber[difference*-1:]
		} else if difference > 0 {
			insertRows(&chamber, difference)
		}

		rowsToInsert := block.FindBottom() + 1
		insertRows(&chamber, rowsToInsert)

		// loop through block movements until rest
		canLower := true
		for canLower {
			direction := string(directions[directionIndex])
			tryShift(chamber, &block, direction)
			canLower = tryShift(chamber, &block, "^")

			// Ensure index
			directionIndex++
			if directionIndex == len(directions) {
				directionIndex = 0
			}
		}

		// write to chamber
		for _, val := range block.Points {
			chamber[val.Y][val.X] = "#"
		}

		// Add completed state
		states = append(states, State{getTowerPositionsState(chamber), directionIndex, i + 1})
	}

	height = len(chamber) - findHighestBlock(chamber)
	return
}

func getTowerPositionsState(chamber [][7]string) [50][7]string {
	towerPositions := [50][7]string{}
	length := int(math.Min(float64(len(towerPositions)), float64(len(chamber))))
	for i := 0; i < length; i++ {
		for j := 0; j < len(towerPositions[i]); j++ {
			towerPositions[i][j] = chamber[i][j]
		}
	}
	return towerPositions
}

func insertRows(chamber *[][7]string, rowCount int) {
	tempSlice := [][7]string{}
	for i := 0; i < rowCount; i++ {
		tempSlice = append(tempSlice, NewRow)
	}
	*chamber = append(tempSlice, *chamber...)
}

func spawnBlock(iterationCount int) Block {
	switch {
	case iterationCount%5 == 0:
		{
			// -
			return Block{[]Point{{0, 2}, {0, 3}, {0, 4}, {0, 5}}}
		}
	case iterationCount%5 == 1:
		{
			// +
			return Block{[]Point{{0, 3}, {1, 3}, {1, 2}, {1, 4}, {2, 3}}}
		}
	case iterationCount%5 == 2:
		{
			// L
			return Block{[]Point{{0, 4}, {1, 4}, {2, 2}, {2, 3}, {2, 4}}}
		}
	case iterationCount%5 == 3:
		{
			// |
			return Block{[]Point{{0, 2}, {1, 2}, {2, 2}, {3, 2}}}
		}
	case iterationCount%5 == 4:
		{
			// Square
			return Block{[]Point{{0, 2}, {1, 2}, {0, 3}, {1, 3}}}
		}
	}
	return Block{}
}

func findHighestBlock(chamber [][7]string) int {
	for i, level := range chamber {
		for _, val := range level {
			if val != SpaceDelimeter {
				return i
			}
		}
	}

	// Is empty
	return len(chamber)
}

// func findHighestColumnBlock(chamber [][7]string, x int) int {
// 	for i := 0; i < len(chamber); i++ {
// 		if chamber[i][x] != SpaceDelimeter {
// 			return i
// 		}
// 	}

// 	return len(chamber)
// }

func tryShift(chamber [][7]string, block *Block, direction string) bool {
	switch direction {
	case "^":
		{
			// Down
			if block.FindBottom() == len(chamber)-1 {
				return false
			}
			for _, point := range block.Points {
				if chamber[point.Y+1][point.X] != SpaceDelimeter {
					return false
				}
			}
			for i := 0; i < len(block.Points); i++ {
				block.Points[i].Y++
			}
			return true
		}
	case "<":
		{
			// Left
			if block.FindLeft() == 0 {
				return false
			}
			for _, point := range block.Points {
				if chamber[point.Y][point.X-1] != SpaceDelimeter {
					return false
				}
			}
			for i := 0; i < len(block.Points); i++ {
				block.Points[i].X--
			}
			return true
		}
	case ">":
		{
			// Right
			if block.FindRight() == len(chamber[0])-1 {
				return false
			}
			for _, point := range block.Points {
				if chamber[point.Y][point.X+1] != SpaceDelimeter {
					return false
				}
			}
			for i := 0; i < len(block.Points); i++ {
				block.Points[i].X++
			}
			return true
		}
	}

	return false
}

func findCycle(states []State) (length int, cycleStart State) {
	var cycleEnd State
	for i := 0; i < len(states)/2; i++ {
		max := int(math.Max(float64(i*2), 1))
		stateOne := states[i]
		stateTwo := states[max]
		if stateOne.Equals(stateTwo) {
			cycleStart = states[i]
			cycleEnd = states[(i * 2)]
			break
		}
	}

	length = cycleEnd.NextIteration - cycleStart.NextIteration

	return
}

func (s State) Equals(s2 State) bool {
	if s.WindIndex != s2.WindIndex {
		return false
	}

	if s.NextIteration%5 != s2.NextIteration%5 {
		return false
	}

	for i := 0; i < len(s.BlockTop); i++ {
		for j := 0; j < len(s.BlockTop[i]); j++ {
			if s.BlockTop[i][j] != s2.BlockTop[i][j] {
				return false
			}
		}
	}

	return true
}
