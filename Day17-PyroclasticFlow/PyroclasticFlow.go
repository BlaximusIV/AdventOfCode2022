package main

import (
	"log"
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
	towerHeight := getSimulatedTowerHeight(string(content), iterations)

	log.Printf("The tower is %d blocks tall \n", towerHeight)

	elapsed := time.Since(startTime)
	log.Printf("Elapsed Time: %s\n", elapsed)
}

func getSimulatedTowerHeight(directions string, iterations int) int {
	// the simulated chamber

	chamber := [][7]string{NewRow, NewRow, NewRow}
	directionIndex := 0
	for i := 0; i < iterations; i++ {
		// Spawn block
		block := spawnBlock(i)

		// Make room
		const spawnMargin = 3
		difference := spawnMargin - findHighestBlock(chamber) // How much space we have

		// Make sure we've got the correct number up front. .. An unnecessary step
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

	}

	return len(chamber) - findHighestBlock(chamber)
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
			// I
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
