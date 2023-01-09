// Logic for building the tower
package main

import (
	"math"
)

const ChamberWidth = 7

// In case I want to switch to another delimeter
const SpaceDelimeter = " "

var NewRow = [7]string{SpaceDelimeter, SpaceDelimeter, SpaceDelimeter, SpaceDelimeter, SpaceDelimeter, SpaceDelimeter, SpaceDelimeter}

func getTowerHeight(directions string, iterations int) int {
	chamber, _ := getChamber(directions, iterations)

	return len(chamber) - findHighestBlock(chamber)
}

func getSimulatedTowerHeight(directions string, iterations int) int {
	// How many iterations are required to find the cycle, so far
	cycleIterations := 3500
	_, states := getChamber(directions, cycleIterations)

	// Find cycle
	// length = number of cycles
	cycleLength, _, _ := findCycle(states)

	// count = number of cycles in requested iterations
	cycleCount := iterations / cycleLength

	// how high is a single cycle
	heightOne := getTowerHeight(directions, cycleLength)
	heightTwo := getTowerHeight(directions, cycleLength*2)
	cycleHeight := heightTwo - heightOne // Should be correct / consistent

	// How many cycles does it take to get to the target?
	remainder := iterations - (cycleCount * cycleLength)

	// Find the tower height of the remainder and add it
	baseTargetHeight := getTowerHeight(directions, cycleLength+remainder)
	totalHeight := baseTargetHeight - cycleHeight + (cycleCount * cycleHeight)

	return totalHeight
}

func getChamber(directions string, iterations int, startingState ...State) (chamber [][7]string, states []State) {
	if len(startingState) > 0 {
		// the simulated chamber
		chamber = startingState[0].BlockTop[:]
	} else {
		chamber = [][7]string{NewRow, NewRow, NewRow}
	}

	states = []State{}
	directionIndex := 0
	for i := 0; i < iterations; i++ {
		block := spawnBlock(i)
		chamber = makeChamberRoom(chamber, i, block)
		lowerBlock(&block, &directionIndex, directions, chamber)

		// write to chamber
		for _, val := range block.Points {
			chamber[val.Y][val.X] = "#"
		}

		// // Add completed state
		states = append(states, State{getTowerPositionsState(chamber), directionIndex, i + 1})
	}

	return
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

func makeChamberRoom(chamber [][7]string, iteration int, block Block) (c [][7]string) {
	const spawnMargin = 3
	// How much space we have
	difference := spawnMargin - findHighestBlock(chamber)

	// Make sure we've got the correct number up front.
	if difference < 0 {
		chamber = chamber[difference*-1:]
	} else if difference > 0 {
		insertRows(&chamber, difference)
	}

	rowsToInsert := block.FindBottom() + 1
	insertRows(&chamber, rowsToInsert)

	return chamber
}

func insertRows(chamber *[][7]string, rowCount int) {
	tempSlice := [][7]string{}
	for i := 0; i < rowCount; i++ {
		tempSlice = append(tempSlice, NewRow)
	}
	*chamber = append(tempSlice, *chamber...)
}

func lowerBlock(block *Block, directionIndex *int, directions string, chamber [][7]string) {
	canLower := true
	for canLower {
		direction := string(directions[*directionIndex])
		tryShift(chamber, block, direction)
		canLower = tryShift(chamber, block, "^")

		// Ensure index
		*directionIndex++
		if *directionIndex == len(directions) {
			*directionIndex = 0
		}
	}
}

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

func getTowerPositionsState(chamber [][7]string) [][7]string {
	stateMaxLength := 50
	length := int(math.Min(float64(stateMaxLength), float64(len(chamber))))

	// towerPositions := [][7]string{}
	towerPositions := chamber[:length]
	return towerPositions
}

func findCycle(states []State) (length int, cycleStart State, cycleEnd State) {
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
