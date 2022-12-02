package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("PuzzleInput.txt")

	if err != nil {
		log.Fatal(err)
	}

	elves := strings.Split(string(content), "\r\n\r\n")

	// Part 1
	max := 0
	for _, elf := range elves {
		sum := 0
		for _, value := range strings.Split(elf, "\r\n") {
			if i, err := strconv.Atoi(value); err == nil {
				sum += i
			}
		}

		if sum > max {
			max = sum
		}
	}

	fmt.Printf("%d Highest calories\n", max)

}
