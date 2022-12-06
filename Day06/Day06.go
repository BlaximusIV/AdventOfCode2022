package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("PuzzleInput.txt")
	if err != nil {
		log.Fatal(err)
	}

	bufferStream := string(content)

	messageUniqueCountRequirement := 14
	foundLock := false
	index := 0
	for !foundLock {
		isUnique := true
		characters := bufferStream[index : index+messageUniqueCountRequirement]
		for i := 0; i < messageUniqueCountRequirement; i++ {
			if strings.Count(characters, string(characters[i])) > 1 {
				isUnique = false
				break
			}
		}

		if isUnique {
			foundLock = true
			break
		}

		index++
	}

	fmt.Printf("Index: %d\n", index+messageUniqueCountRequirement)
}
