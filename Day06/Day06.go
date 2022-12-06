/* This solution neglects a few error prevention techniques because the data is known ahead of time*/
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

	uniqueRequirement := 14
	index := 0
	for {
		isUnique := true
		characters := bufferStream[index : index+uniqueRequirement]
		for _, char := range characters {
			if strings.Count(characters, string(char)) > 1 {
				isUnique = false
				break
			}
		}

		if isUnique {
			break
		}

		index++
	}

	fmt.Printf("Processed characters: %d\n", index+uniqueRequirement)
}
