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

func (c Cube) CountExposed(drop Droplet, excludeBubbles bool) int {
	borderedSides := 0
	borders := c.GetNeighbors()

	for i := range borders {
		if drop.Cubes[borders[i]] {
			borderedSides++
		} else if excludeBubbles && borders[i].IsBubbleCoord(drop) {
			borderedSides++
		}
	}

	return 6 - borderedSides
}

func (c Cube) IsBubbleCoord(drop Droplet) bool {
	// If there is a block in all six directions
	// This assumes no hooks (J) or negative coordinates, and seems to be correct
	// x, false if goes to max or min
	x := c.x + 1
	for x < drop.MaxX {
		if drop.Cubes[Cube{x, c.y, c.z}] {
			break
		}
		x++
	}
	if x >= drop.MaxX {
		return false
	}
	nx := c.x - 1
	for nx > 0 {
		if drop.Cubes[Cube{nx, c.y, c.z}] {
			break
		}
		nx--
	}
	if nx <= 0 {
		return false
	}
	y := c.y + 1
	for y < drop.MaxY {
		if drop.Cubes[Cube{c.x, y, c.z}] {
			break
		}
		y++
	}
	if y >= drop.MaxY {
		return false
	}
	ny := c.y - 1
	for y > 0 {
		if drop.Cubes[Cube{c.x, ny, c.z}] {
			break
		}
		y--
	}
	if ny <= 0 {
		return false
	}
	z := c.z + 1
	for z <= drop.MaxZ {
		if drop.Cubes[Cube{c.x, c.y, z}] {
			break
		}
		z++
	}
	if z >= drop.MaxZ {
		return false
	}
	nz := c.z - 1
	for nz >= 0 {
		if drop.Cubes[Cube{c.x, c.y, nz}] {
			break
		}
		nz--
	}
	if nz <= 0 {
		return false
	}

	return true
}

func main() {
	startTime := time.Now()

	content, _ := os.ReadFile("TestInput.txt")

	drop := getDrop(string(content))

	// Part 1
	sum := 0
	for c := range drop.Cubes {
		sum += c.CountExposed(drop, false)
	}
	log.Printf("Surface Area: %d", sum)

	// Part 2
	sum = 0
	for c := range drop.Cubes {
		sum += c.CountExposed(drop, true)
	}

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
