package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()

	content, _ := os.ReadFile("PuzzleInput.txt")
	pairs := strings.Split(string(content), "\r\n\r\n")

	// Part 1
	correctSum := getCorrectOrderedPairsSum(pairs)
	log.Printf("Sum of pair indexes in correct order: %d", correctSum)

	elapsed := time.Since(startTime)
	log.Printf("Elapsed Time: %s\n", &elapsed)
}

func getCorrectOrderedPairsSum(pairs []string) (sum int) {
	for i, pairBlock := range pairs {
		pair := strings.Split(pairBlock, "\r\n")
		if isCorrectOrder(pair[0], pair[1]) {
			sum += i + 1
		}
	}

	return
}

func isCorrectOrder(lhs string, rhs string) bool {
	index := 0
	for {
		lchar, rchar := string(lhs[index]), string(rhs[index])

		// If they're the same
		if lchar == rchar {
			index++
			continue
		}

		// They're both numbers
		lhNum, lerr := strconv.Atoi(string(lchar))
		rhNum, rerr := strconv.Atoi(string(rchar))
		if lerr == nil && rerr == nil {
			return lhNum < rhNum
		}

		// "[", ","
		if lchar == "]" || rchar == "]" {
			return lchar == "]"
		}

		// One is a number
		if rchar == "[" {
			lhs = lhs[:index] + "[" + string(lhs[index]) + "]" + lhs[index+1:]
			index++
			continue
		} else if lchar == "[" {
			rhs = rhs[:index] + "[" + string(rhs[index]) + "]" + rhs[index+1:]
			index++
			continue
		} else if lerr == nil {
			return false
		} else {
			return true
		}

	}
}
