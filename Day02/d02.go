package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const TIE string = "TIE"
const WIN string = "WIN"
const LOSS string = "LOSS"
const A string = "A"
const B string = "B"
const C string = "C"
const X string = "X"
const Y string = "Y"
const Z string = "Z"

var PlayScores = map[string]int{
	X: 1,
	Y: 2,
	Z: 3,
}

var OutcomeScores = map[string]int{
	LOSS: 0,
	TIE:  3,
	WIN:  6,
}

func main() {
	content, err := ioutil.ReadFile("PuzzleInput.txt")

	if err != nil {
		log.Fatal(err)
	}

	rounds := strings.Split(string(content), "\r\n")

	var roundScores []int
	for _, round := range rounds {
		roundScores = append(roundScores, calculateRoundScore(round))
	}

	// Part 1
	sum := 0
	for _, score := range roundScores {
		sum += score
	}

	fmt.Printf("Total Score: %d \n", sum)
}

func calculateRoundScore(round string) (score int) {
	plays := strings.Split(round, " ")

	outcome := getOutcome(plays[0], plays[1])
	score = PlayScores[plays[1]] + OutcomeScores[outcome]

	return
}

func getOutcome(oppPlay string, myPlay string) (outcome string) {
	switch myPlay {
	case X:
		switch oppPlay {
		case A:
			outcome = TIE
		case B:
			outcome = LOSS
		default:
			outcome = WIN
		}
	case Y:
		switch oppPlay {
		case B:
			outcome = TIE
		case C:
			outcome = LOSS
		default:
			outcome = WIN
		}
	case Z:
		switch oppPlay {
		case C:
			outcome = TIE
		case A:
			outcome = LOSS
		default:
			outcome = WIN
		}
	}

	return
}
