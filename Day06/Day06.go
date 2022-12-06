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
	foundLock := false
	index := 0
	for !foundLock {
		characters := bufferStream[index : index+4]

		if strings.Count(characters, string(characters[0])) == 1 &&
			strings.Count(characters, string(characters[1])) == 1 &&
			strings.Count(characters, string(characters[2])) == 1 &&
			strings.Count(characters, string(characters[3])) == 1 {
			foundLock = true
			continue
		}

		index++
	}

	fmt.Printf("Index: %d\n", index+4)
}
