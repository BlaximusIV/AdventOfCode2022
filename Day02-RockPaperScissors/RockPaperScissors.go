package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	TIE  string = "TIE"
	WIN  string = "WIN"
	LOSS string = "LOSS"
	A    string = "A"
	B    string = "B"
	C    string = "C"
	X    string = "X"
	Y    string = "Y"
	Z    string = "Z"
)

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

var RequiredOutcome = map[string]string{
	X: LOSS,
	Y: TIE,
	Z: WIN,
}

var Plays = map[string]string{
	A: X,
	B: Y,
	C: Z,
}

func main() {
	content, err := ioutil.ReadFile("PuzzleInput.txt")

	if err != nil {
		log.Fatal(err)
	}

	rounds := strings.Split(string(content), "\r\n")

	var roundScores []int
	var requiredRoundScores []int
	for _, round := range rounds {
		roundScores = append(roundScores, calculateRoundScore(round))
		requiredRoundScores = append(requiredRoundScores, calculateRequiredRoundScore(round))
	}

	// Part 1
	sum := 0
	for _, score := range roundScores {
		sum += score
	}

	fmt.Printf("Total Score: %d \n", sum)

	// Part 2
	requiredSum := 0
	for _, score := range requiredRoundScores {
		requiredSum += score
	}

	fmt.Printf("Total Required Score: %d \n", requiredSum)
}

func calculateRoundScore(round string) (score int) {
	plays := strings.Split(round, " ")

	outcome := getOutcome(plays[0], plays[1])
	score = PlayScores[plays[1]] + OutcomeScores[outcome]

	return
}

func getOutcome(oppPlay string, myPlay string) (outcome string) {
	if Plays[oppPlay] == myPlay {
		outcome = TIE
		return
	}

	switch myPlay {
	case X:
		if oppPlay == B {
			outcome = LOSS
		}
	case Y:
		if oppPlay == C {
			outcome = LOSS
		}
	case Z:
		if oppPlay == A {
			outcome = LOSS
		}
	}

	if outcome == "" {
		outcome = WIN
	}

	return
}

func calculateRequiredRoundScore(round string) (score int) {
	plan := strings.Split(round, " ")

	oppPlay := plan[0]
	requiredOutcome := RequiredOutcome[plan[1]]

	myPlay := getRequiredPlay(oppPlay, requiredOutcome)

	score = PlayScores[myPlay] + OutcomeScores[requiredOutcome]

	return
}

func getRequiredPlay(oppPlay string, requiredOutcome string) (play string) {
	switch requiredOutcome {
	case WIN:
		switch oppPlay {
		case A:
			play = Y
		case B:
			play = Z
		case C:
			play = X
		}
	case LOSS:
		switch oppPlay {
		case A:
			play = Z
		case B:
			play = X
		case C:
			play = Y
		}
	case TIE:
		play = Plays[oppPlay]
	}

	return
}
