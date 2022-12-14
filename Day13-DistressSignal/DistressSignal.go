/*Gross solution. Needs refactoring.*/
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

	// Part 2
	lines := getPackets(pairs)
	lines = append(lines, "[[2]]", "[[6]]")

	sort(lines)
	decoderKey := getDecoderKey(lines)

	log.Printf("Decoder Key: %d", decoderKey)

	elapsed := time.Since(startTime)
	log.Printf("Elapsed Time: %s\n", &elapsed)
}

func getPackets(pairs []string) (packets []string) {
	packets = []string{}
	for _, pair := range pairs {
		packets = append(packets, strings.Split(pair, "\r\n")...)
	}
	return
}

func sort(packets []string) {
	// The dataset is small, bubble sort it is!
	for i := 0; i < len(packets)-1; i++ {
		for j := i + 1; j < len(packets); j++ {
			if !isCorrectOrder(packets[i], packets[j]) {
				temp := packets[i]
				packets[i] = packets[j]
				packets[j] = temp
			}
		}
	}
}

func getDecoderKey(packets []string) int {
	dividerPackets := [2]int{}

	for i, val := range packets {
		if val == "[[2]]" {
			dividerPackets[0] = i
		} else if val == "[[6]]" {
			dividerPackets[1] = i
		}
	}

	return (dividerPackets[0] + 1) * (dividerPackets[1] + 1)
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

// Ugly as sin brute force comparison, but it works
func isCorrectOrder(lhs string, rhs string) bool {
	// Separate index
	lhIndex := 0
	rhIndex := 0
	for {
		lchar, rchar := lhs[lhIndex], rhs[rhIndex]
		lstr, rstr := string(lchar), string(rchar)

		lisNum := utf8Numbers.Contains(lchar)
		risNum := utf8Numbers.Contains(rchar)

		if lisNum {
			for utf8Numbers.Contains(lhs[lhIndex+1]) {
				lstr += string(lhs[lhIndex+1])
				lhIndex++
			}

		}

		if risNum {
			for utf8Numbers.Contains(rhs[rhIndex+1]) {
				rstr += string(rhs[rhIndex+1])
				rhIndex++
			}
		}

		areBothNumbers := lisNum && risNum
		if lchar == rchar && !areBothNumbers {
			lhIndex++
			rhIndex++
			continue
		}

		// They're both numbers
		if areBothNumbers {

			lNum, _ := strconv.Atoi(lstr)
			rNum, _ := strconv.Atoi(rstr)

			if lNum == rNum {
				atEnd := lhIndex+1 >= len(lhs) || rhIndex+1 >= len(rhs)
				if !atEnd {
					lhIndex++
					rhIndex++
				}
				continue
			}

			return lNum < rNum
		}

		// '[', ','
		if lchar == ']' || rchar == ']' {
			return lchar == ']'
		}

		// One is a number
		if rchar == '[' {
			diff := len(lstr) - 1
			lhs = lhs[:lhIndex-diff] + "[" + lstr + "]" + lhs[lhIndex+1:]
			lhIndex += 1 - diff
			rhIndex++
			continue
		} else if lchar == '[' {
			diff := len(rstr) - 1
			rhs = rhs[:rhIndex-diff] + "[" + rstr + "]" + rhs[rhIndex+1:]
			lhIndex++
			rhIndex += 1 - diff
			continue
		} else if lisNum {
			return false
		} else {
			return true
		}

	}
}
