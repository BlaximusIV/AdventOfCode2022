package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()

	input, _ := os.ReadFile("Input.txt")

	elves := getElves(string(input))
	grid := Grid{}
	grid.MakeGrid(elves)

	// Part 1
	diffuseXTimes(&grid, 10)
	fmt.Printf("Empty tiles after 10 rounds: %d\n", grid.CountEmptyTiles())

	// Part 2
	grid = Grid{}
	grid.MakeGrid(elves)
	diffusionCount := getDiffusionRoundCount(&grid)
	fmt.Printf("Required rounds to full diffusion: %d\n", diffusionCount)

	elapsed := time.Since(startTime)
	fmt.Printf("Elapsed time: %v\n", elapsed)
}

func getElves(input string) []Elf {
	elves := []Elf{}
	for i, line := range strings.Split(input, "\n") {
		for j, char := range strings.Split(line, "") {
			if char == "#" {
				elves = append(elves, Elf{i, j})
			}
		}
	}
	return elves
}

func getDiffusionRoundCount(g *Grid) int {
	roundCount := 1
	direction := North
	isDiffusing := true
	for isDiffusing {
		isDiffusing = diffuseElves(g, direction)

		if isDiffusing {
			roundCount++
		}

		direction = direction.Next()
	}

	return roundCount
}

func diffuseXTimes(g *Grid, times int) {
	direction := North
	for i := 0; i < 10; i++ {
		diffuseElves(g, direction)
		direction = direction.Next()
	}
}

func diffuseElves(g *Grid, startDirection Direction) (diffused bool) {
	proposalMap := map[Elf]int{}

	proposals := []Proposal{}
	for _, elf := range g.Elves {
		// Check if there are any around
		if !g.Map.HasAnyNeighbors(elf) {
			proposals = append(proposals, Proposal{elf, elf})
		} else {
			proposal := findProposal(g.Map, startDirection, elf)
			proposalMap[proposal.To] += 1
			proposals = append(proposals, proposal)
		}
	}

	finalElves := []Elf{}
	// For each elf proposing a move
	for _, p := range proposals {
		// if unique location, move
		if proposalMap[p.To] == 1 {
			g.Move(p.From, p.To)

			finalElves = append(finalElves, p.To)
			diffused = true
		} else {
			finalElves = append(finalElves, p.From)
		}
	}

	g.Elves = finalElves
	return
}

func findProposal(m Map, direction Direction, e Elf) Proposal {
	const CardinalDirections = 4

	d := direction
	for i := CardinalDirections; i > 0; i-- {
		if !m.HasNeighbors(e, d) {
			newCoord := getNewCoordinate(d, e)
			return Proposal{e, newCoord}
		}

		d = d.Next()
	}

	return Proposal{e, e}
}

func getNewCoordinate(d Direction, e Elf) Elf {
	switch d {
	case North:
		{
			return Elf{e.Y - 1, e.X}
		}
	case South:
		{
			return Elf{e.Y + 1, e.X}
		}
	case West:
		{
			return Elf{e.Y, e.X - 1}
		}
	}

	return Elf{e.Y, e.X + 1}
}
