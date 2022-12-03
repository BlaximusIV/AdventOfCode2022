package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("PuzzleInput.txt")

	if err != nil {
		log.Fatal(err)
	}

	rucksacks := strings.Split(string(content), "\r\n")

	sum := 0
	for _, sack := range rucksacks {
		sum += getSackDuplicatePriority(sack)
	}

	fmt.Printf("Priorities sum: %d", sum)
}

func getSackDuplicatePriority(sack string) (priority int) {
	// Find duplicate letter
	duplicateItem := findDuplicateItem(sack)

	char := []rune(duplicateItem)

	// Find letter value
	priority = findLetterValue(char[0])

	return
}

func findDuplicateItem(sack string) (duplicate string) {
	pocketSize := len(sack) / 2
	pocketOne := sack[:pocketSize]
	pocketTwo := sack[pocketSize:]

	for _, item := range strings.Split(pocketTwo, "") {
		if strings.Contains(pocketOne, item) {
			duplicate = item
			return
		}
	}

	return
}

func findLetterValue(item rune) (value int) {
	asciiLowerOffset := 96
	asciiUperOffset := 38
	value = int(item)

	if value > 95 {
		value -= asciiLowerOffset
	} else {
		value -= asciiUperOffset
	}

	return
}
