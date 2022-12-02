package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile("PuzzleInput.txt")

	if err != nil {
		log.Fatal(err)
	}

	elves := strings.Split(string(content), "\r\n\r\n")

	var elfCalories []int
	for _, elf := range elves {
		elfCalories = append(elfCalories, getElfSum(elf))
	}

	sort.Slice(elfCalories, func(i, j int) bool {
		return elfCalories[i] > elfCalories[j]
	})

	// Part 1
	fmt.Printf("%d Highest calories\n", elfCalories[0])

	// Part 2
	topthree := elfCalories[:3]
	sum := 0
	for _, val := range topthree {
		sum += val
	}

	fmt.Printf("Sum of top 3: %d\n", sum)
}

func getElfSum(elf string) (sum int) {
	for _, value := range strings.Split(elf, "\r\n") {
		if i, err := strconv.Atoi(value); err == nil {
			sum += i
		}
	}

	return
}
