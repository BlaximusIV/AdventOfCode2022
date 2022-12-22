package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Droplet struct {
	Cubes            map[Cube]bool
	MaxX, MaxY, MaxZ int
}

type Cube struct {
	x, y, z int
}

func (c Cube) GetNeighbors() []Cube {
	return []Cube{{c.x + 1, c.y, c.z}, {c.x - 1, c.y, c.z}, {c.x, c.y + 1, c.z}, {c.x, c.y - 1, c.z}, {c.x, c.y, c.z + 1}, {c.x, c.y, c.z - 1}}
}

func (c Cube) CountExposed(drop Droplet) int {
	borderedSides := 0
	borders := c.GetNeighbors()

	for i := range borders {
		if drop.Cubes[borders[i]] {
			borderedSides++
		}
	}

	return 6 - borderedSides
}

func main() {
	startTime := time.Now()

	content, _ := os.ReadFile("Input.txt")

	drop := getDrop(string(content))

	// Part 1
	sum := 0
	for c := range drop.Cubes {
		sum += c.CountExposed(drop)
	}
	log.Printf("Surface Area: %d", sum)

	// Part 2
	sum = getFloodObstructionCount(drop)

	log.Printf("Surface Area: %d", sum)

	elapsed := time.Since(startTime)
	log.Printf("Elapsed Time: %s\n", elapsed)
}

func getDrop(input string) Droplet {
	cubes := map[Cube]bool{}
	maxx, maxy, maxz := 0, 0, 0
	for _, val := range strings.Split(input, "\r\n") {
		xyz := strings.Split(val, ",")
		x, _ := strconv.Atoi(xyz[0])
		y, _ := strconv.Atoi(xyz[1])
		z, _ := strconv.Atoi(xyz[2])
		if x > maxx {
			maxx = x
		}
		if y > maxy {
			maxy = y
		}
		if z > maxz {
			maxz = z
		}
		cubes[Cube{x, y, z}] = true
	}
	return Droplet{cubes, maxx, maxy, maxz}
}

// Basically BFS
func getFloodObstructionCount(drop Droplet) int {
	visited := map[Cube]bool{}
	frontier := []Cube{{-1, -1, -1}}
	obstructionCount := 0
	for len(frontier) > 0 {
		cube := frontier[0]
		visited[cube] = true
		frontier = frontier[1:]

		neighbors := cube.GetNeighbors()

		inBoundNeighbors := []Cube{}
		for _, val := range neighbors {
			if val.x < -1 || val.x > drop.MaxX+1 {
				continue
			} else if val.y < -1 || val.y > drop.MaxY+1 {
				continue
			} else if val.z < -1 || val.z > drop.MaxZ+1 {
				continue
			}

			inBoundNeighbors = append(inBoundNeighbors, val)
		}

		nonQueuedNeighbors := []Cube{}
		for _, val := range inBoundNeighbors {
			contains := false
			for _, frontierVal := range frontier {
				if val == frontierVal {
					contains = true
					break
				}
			}
			if !contains {
				nonQueuedNeighbors = append(nonQueuedNeighbors, val)
			}
		}

		nonBlockedNeighbors := []Cube{}
		for _, val := range nonQueuedNeighbors {
			if drop.Cubes[val] {
				// entry from this direction is unavailable. Counts as a single face of a single lava drop block
				obstructionCount++
			} else {
				nonBlockedNeighbors = append(nonBlockedNeighbors, val)
			}
		}

		nonVisitedNeighbors := []Cube{}
		for _, val := range nonBlockedNeighbors {
			if !visited[val] {
				nonVisitedNeighbors = append(nonVisitedNeighbors, val)
			}
		}

		frontier = append(frontier, nonVisitedNeighbors...)
	}

	return obstructionCount
}
