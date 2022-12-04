/*
	A simpler way to do this probably would have been to create slices and use the built-in functionality to see if one slice contained another
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type ZonePairAssignments struct {
	AssignOneLow, AssignOneHigh, AssignTwoLow, AssignTwoHigh int
}

func main() {
	content, err := ioutil.ReadFile("PuzzleInput.txt")

	if err != nil {
		log.Fatal(err)
	}

	zonePairs := strings.Split(string(content), "\r\n")

	redundantZonePairCount := 0
	overlappingZonePairCount := 0
	for _, pair := range zonePairs {
		assignments := getZonePairAssignments(pair)
		if isRedundantZonePair(assignments) {
			redundantZonePairCount++
		}

		if isOverlappingZonePair(assignments) {
			overlappingZonePairCount++
		}
	}

	// Part 1
	fmt.Printf("%d redundant zone pairs\n", redundantZonePairCount)

	// Part 2
	fmt.Printf("%d overlapping zone pairs\n", overlappingZonePairCount)
}

func getZonePairAssignments(zonePair string) ZonePairAssignments {
	assignments := strings.Split(zonePair, ",")
	assignmentOne := strings.Split(assignments[0], "-")
	assignmentTwo := strings.Split(assignments[1], "-")

	assignOneLow, errOneLow := strconv.Atoi(assignmentOne[0])
	assignOneHigh, errOneHigh := strconv.Atoi(assignmentOne[1])
	assignTwoLow, errTwoLow := strconv.Atoi(assignmentTwo[0])
	assignTwoHigh, errTwoHigh := strconv.Atoi(assignmentTwo[1])

	if errOneHigh != nil || errOneLow != nil || errTwoHigh != nil || errTwoLow != nil {
		log.Fatal("Unable to parse assigned zones")
	}

	return ZonePairAssignments{assignOneLow, assignOneHigh, assignTwoLow, assignTwoHigh}
}

func isRedundantZonePair(assignments ZonePairAssignments) (isRedundant bool) {
	assignmentOneEnvelops := assignments.AssignOneLow <= assignments.AssignTwoLow && assignments.AssignOneHigh >= assignments.AssignTwoHigh
	assignmentTwoEnvelops := assignments.AssignTwoLow <= assignments.AssignOneLow && assignments.AssignTwoHigh >= assignments.AssignOneHigh

	if assignmentOneEnvelops || assignmentTwoEnvelops {
		isRedundant = true
	}

	return
}

func isOverlappingZonePair(assignments ZonePairAssignments) (isOverlapping bool) {
	oneHasEnvelopedLow := assignments.AssignOneLow >= assignments.AssignTwoLow && assignments.AssignOneLow <= assignments.AssignTwoHigh
	oneHasEnvelopedHigh := assignments.AssignOneHigh >= assignments.AssignTwoLow && assignments.AssignOneHigh <= assignments.AssignTwoHigh

	twoHasEnvelopedLow := assignments.AssignTwoLow >= assignments.AssignOneLow && assignments.AssignTwoLow <= assignments.AssignOneHigh
	twoHasEnvelopedHigh := assignments.AssignTwoHigh >= assignments.AssignOneLow && assignments.AssignTwoHigh <= assignments.AssignOneHigh

	isOverlapping = oneHasEnvelopedLow || oneHasEnvelopedHigh || twoHasEnvelopedLow || twoHasEnvelopedHigh

	return
}
