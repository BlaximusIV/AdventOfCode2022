package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	content, _ := os.ReadFile("PuzzleInput.txt")

	treeMap := importTreeMap(string(content))

	visibleTreeCount, highestViewRating := evaluateVisibleTrees(treeMap)

	// Part 1
	log.Printf("Visible tree count: %d\n", visibleTreeCount)

	// Part 2
	log.Printf("Highest tree view rating: %d\n", highestViewRating)
	elapsed := time.Since(start)
	log.Printf("Elapsed Time: %s\n", elapsed)
}

func importTreeMap(input string) (treeMap [][]int) {
	for _, row := range strings.Split(input, "\r\n") {
		treeSizeRow := []int{}
		for _, sizeString := range strings.Split(row, "") {
			size, _ := strconv.Atoi(sizeString)
			treeSizeRow = append(treeSizeRow, size)
		}

		treeMap = append(treeMap, treeSizeRow)
	}

	return
}

func evaluateVisibleTrees(treeMap [][]int) (visibleTreeCount int, highestViewRating int) {
	// Grid is a square of gridSize x gridSize
	gridSize := len(treeMap)
	for y := range treeMap {
		for x := range treeMap[y] {
			isVisibleFromEdge, totalViewRating := getTreeView(x, y, gridSize, treeMap)

			if isVisibleFromEdge {
				visibleTreeCount++
			}

			if highestViewRating < totalViewRating {
				highestViewRating = totalViewRating
			}
		}
	}

	return
}

func getTreeView(x int, y int, gridSize int, treeMap [][]int) (isVisibleFromEdge bool, totalViewRating int) {
	isTallestNorth, northViewRating := getTreeNorthView(x, y, gridSize, treeMap)
	isTallestWest, westViewRating := getTreeWestView(x, y, gridSize, treeMap)
	isTallestEast, eastViewRating := getTreeEastView(x, y, gridSize, treeMap)
	isTallestSouth, southViewRating := getTreeSouthView(x, y, gridSize, treeMap)

	isVisibleFromEdge = isTallestNorth || isTallestWest || isTallestEast || isTallestSouth

	totalViewRating = northViewRating * westViewRating * eastViewRating * southViewRating

	return
}

func getTreeNorthView(x int, y int, gridSize int, treeMap [][]int) (isTallestFromNorth bool, northViewRating int) {
	coordinateTreeSize := treeMap[y][x]
	isTallestFromNorth = true
	for n := y - 1; n >= 0; n-- {
		northViewRating++
		if treeMap[n][x] >= coordinateTreeSize {
			isTallestFromNorth = false
			break
		}
	}

	return
}

func getTreeWestView(x int, y int, gridSize int, treeMap [][]int) (isTallestFromWest bool, westViewRating int) {
	coordinateTreeSize := treeMap[y][x]
	isTallestFromWest = true
	for w := x - 1; w >= 0; w-- {
		westViewRating++
		if treeMap[y][w] >= coordinateTreeSize {
			isTallestFromWest = false
			break
		}
	}

	return
}

func getTreeEastView(x int, y int, gridSize int, treeMap [][]int) (isTallestFromEast bool, eastViewRating int) {
	coordinateTreeSize := treeMap[y][x]
	isTallestFromEast = true
	for e := x + 1; e < gridSize; e++ {
		eastViewRating++
		if treeMap[y][e] >= coordinateTreeSize {
			isTallestFromEast = false
			break
		}
	}

	return
}

func getTreeSouthView(x int, y int, gridSize int, treeMap [][]int) (isTallestFromSouth bool, southViewRating int) {
	coordinateTreeSize := treeMap[y][x]
	isTallestFromSouth = true
	for s := y + 1; s < gridSize; s++ {
		southViewRating++
		if treeMap[s][x] >= coordinateTreeSize {
			isTallestFromSouth = false
			break
		}
	}

	return
}
