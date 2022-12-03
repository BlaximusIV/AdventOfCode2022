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

	// Part 1
	sum := 0
	for _, sack := range rucksacks {
		sum += getDuplicatePocketItemPriority(sack)
	}

	fmt.Printf("Priorities sum: %d\n", sum)

	// Part 2
	badgeSum := 0
	for i := 0; i < len(rucksacks); i += 3 {
		group := rucksacks[i : i+3]
		badgeSum += getSacksBadgePriority(group)
	}

	fmt.Printf("Badge priorities sum: %d\n", badgeSum)
}

func getDuplicatePocketItemPriority(sack string) (priority int) {
	// Find duplicate letter
	duplicateItem := findDuplicatePocketItem(sack)

	// Find letter value
	priority = findLetterValue(duplicateItem)

	return
}

func findDuplicatePocketItem(sack string) (duplicate rune) {
	pocketSize := len(sack) / 2
	pocketOne := sack[:pocketSize]
	pocketTwo := sack[pocketSize:]

	for _, item := range pocketTwo {
		if strings.Contains(pocketOne, string(item)) {
			duplicate = item
			return
		}
	}

	return
}

func findLetterValue(item rune) (value int) {
	asciiLowerOffset := 96
	asciiUpperOffset := 38
	value = int(item)

	if value > 95 {
		value -= asciiLowerOffset
	} else {
		value -= asciiUperOffset
	}

	return
}

func getSacksBadgePriority(sacks []string) (priority int) {
	duplicateItem := findDuplicateSacksItem(sacks)

	priority = findLetterValue(duplicateItem)

	return
}

func findDuplicateSacksItem(sacks []string) (duplicate rune) {
	for _, item := range sacks[2] {
		oneHas := strings.Contains(sacks[0], string(item))
		twoHas := strings.Contains(sacks[1], string(item))

		if oneHas && twoHas {
			// Then we know the item is the badge
			duplicate = item
			return
		}
	}

	return
}
