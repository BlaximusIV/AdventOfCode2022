package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Utf8Numbers struct {
	Values []int
}

func (a Utf8Numbers) Contains(char byte) bool {
	for _, val := range a.Values {
		if val == int(char) {
			return true
		}
	}
	return false
}

// The utf8 codes representing the numbers 0-9. Directly comparable to bytes
var utf8Numbers = Utf8Numbers{[]int{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}}

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
	// Separate index
	lhIndex := 0
	rhIndex := 0
	for {
		// TODO: Get numbers up front
		lchar, rchar := lhs[lhIndex], rhs[rhIndex]
		lstr, rstr := string(lchar), string(rchar)
		// If they're the same
		if lchar == rchar {
			lhIndex++
			rhIndex++
			continue
		}

		// They're both numbers
		lisNum := utf8Numbers.Contains(lchar)
		risNum := utf8Numbers.Contains(rchar)
		if lisNum && risNum {
			lNumString := ""
			rNumString := ""

			// Get whole numbers
			for utf8Numbers.Contains(lhs[lhIndex]) {
				lNumString += string(lhs[lhIndex])
				lhIndex++
			}

			for utf8Numbers.Contains(rhs[rhIndex]) {
				rNumString += string(rhs[rhIndex])
				rhIndex++
			}

			lNum, _ := strconv.Atoi(lNumString)
			rNum, _ := strconv.Atoi(rNumString)

			return lNum < rNum
		}

		// '[', ','
		if lchar == ']' || rchar == ']' {
			return lchar == ']'
		}

		// One is a number
		if rchar == '[' {
			lhs = lhs[:lhIndex] + "[" + lstr + "]" + lhs[lhIndex+1:]
			lhIndex++
			rhIndex++
			continue
		} else if lchar == '[' {
			rhs = rhs[:rhIndex] + "[" + rstr + "]" + rhs[rhIndex+1:]
			lhIndex++
			rhIndex++
			continue
		} else if lisNum {
			return false
		} else {
			return true
		}

	}
}
