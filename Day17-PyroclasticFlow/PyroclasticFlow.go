package main

import (
	"log"
	"os"
	"time"
)

func main() {
	startTime := time.Now()

	content, _ := os.ReadFile("Input.txt")

	// Part1
	iterations := 2022
	towerHeight := getTowerHeight(string(content), iterations)
	log.Printf("The tower is %d blocks tall \n", towerHeight)

	// // Part2
	iterations = 1_000_000_000_000
	towerHeight = getSimulatedTowerHeight(string(content), iterations)
	log.Printf("The simulated tower is %d blocks tall \n", towerHeight)

	elapsed := time.Since(startTime)
	log.Printf("Elapsed Time: %s\n", elapsed)
}
