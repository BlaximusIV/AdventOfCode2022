package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Cube struct {
	x, y, z int
}

func (c Cube) CountExposed(cubes map[Cube]bool) int {
	borderedSides := 0
	borders := []Cube{{c.x + 1, c.y, c.z}, {c.x - 1, c.y, c.z}, {c.x, c.y + 1, c.z}, {c.x, c.y - 1, c.z}, {c.x, c.y, c.z + 1}, {c.x, c.y, c.z - 1}}

	for i := range borders {
		if cubes[borders[i]] {
			borderedSides++
		}
	}

	return 6 - borderedSides
}

func main() {
	startTime := time.Now()

	content, _ := os.ReadFile("Input.txt")

	cubes := getCubes(string(content))

	// Part 1
	sum := 0
	for c := range cubes {
		sum += c.CountExposed(cubes)
	}

	log.Printf("Surface Area: %d", sum)

	elapsed := time.Since(startTime)
	log.Printf("Elapsed Time: %s\n", elapsed)
}

func getCubes(input string) map[Cube]bool {
	cubes := map[Cube]bool{}
	for _, val := range strings.Split(input, "\r\n") {
		xyz := strings.Split(val, ",")
		x, _ := strconv.Atoi(xyz[0])
		y, _ := strconv.Atoi(xyz[1])
		z, _ := strconv.Atoi(xyz[2])
		cubes[Cube{x, y, z}] = true
	}
	return cubes
}
