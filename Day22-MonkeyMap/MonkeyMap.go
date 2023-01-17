package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()

	input, _ := os.ReadFile("Input.txt")

	terrain, directions := parse2DInput(string(input))

	// Part1
	password := get2DPassword(terrain, directions)

	fmt.Printf("2D Password: %d\n", password)

	elapsed := time.Since(startTime)
	fmt.Printf("Elapsed time: %v\n", elapsed)
}

func parse2DInput(input string) (MonkeyMap, []Direction) {
	parts := strings.Split(input, "\n\n")

	plot := MonkeyMap{}
	for _, line := range strings.Split(parts[0], "\n") {
		row := []string{}
		row = append(row, strings.Split(line, "")...)
		plot = append(plot, row)
	}

	magExp, _ := regexp.Compile(`(\d+)`)
	turnExp, _ := regexp.Compile(`([RL])`)

	magnitudes := magExp.FindAllString(parts[1], -1)
	turns := turnExp.FindAllString(parts[1], -1)

	directions := []Direction{}
	for i, mag := range magnitudes {
		val, _ := strconv.Atoi(mag)
		turn := "None"
		if len(turns) > i {
			turn = turns[i]
		}
		directions = append(directions, Direction{val, turn})
	}

	return plot, directions
}

func get2DPassword(terrain MonkeyMap, directions []Direction) int {
	const RowMult = 1000
	const ColMult = 4

	startCoord := terrain.leftMostCoordinate(0)
	cursor := Cursor{startCoord, Right}

	followDirections(terrain, directions, &cursor)

	// Column / rows start at 1 for scoring
	rowScore := RowMult * (cursor.Position.Y + 1)
	colScore := ColMult * (cursor.Position.X + 1)
	faceScore := int(cursor.Facing)
	return rowScore + colScore + faceScore
}

func followDirections(terrain MonkeyMap, directions []Direction, cursor *Cursor) {
	// visited := map[Coordinate]bool{{cursor.Position.Y, cursor.Position.X}: true}
	for _, direction := range directions {
		// Try to move direction facing
		for i := 0; i < direction.Magnitude; i++ {
			tryMove(terrain, cursor)
			// visited[cursor.Position] = true
		}

		turn(cursor, direction.TurnDirection)
		// print(direction, terrain, visited)
	}
}

func tryMove(terrain MonkeyMap, cursor *Cursor) {
	coord := cursor.NextCoord()

	// If invalid, adjust
	switch cursor.Facing {
	case Right:
		{
			if coord.X >= len(terrain[coord.Y]) {
				coord = terrain.leftMostCoordinate(coord.Y)
			}
		}
	case Down:
		{
			if coord.Y >= len(terrain) || coord.X >= len(terrain[coord.Y]) {
				coord = terrain.topMostCoordinate(coord.X)
			}
		}
	case Left:
		{
			if coord.X < 0 || terrain[coord.Y][coord.X] == " " {
				coord = terrain.rightMostCoordinate(coord.Y)
			}
		}
	case Up:
		{
			if coord.Y < 0 || terrain[coord.Y][coord.X] == " " {
				coord = terrain.botMostCoordinate(coord.Y, coord.X)
			}
		}
	}

	// If blocked return
	if terrain[coord.Y][coord.X] == "#" {
		return
	}

	// else update cursor Coordinate
	cursor.Position = coord
}

// Facing is enum, so can use ++, --
func turn(cursor *Cursor, direction string) {
	// R/ight L/eft
	if direction == "R" {
		if cursor.Facing == Up {
			cursor.Facing = Right
		} else {
			cursor.Facing++
		}
	} else if direction == "L" {
		if cursor.Facing == Right {
			cursor.Facing = Up
		} else {
			cursor.Facing--
		}
	}
}

// func print(direction Direction, terrain MonkeyMap, visited map[Coordinate]bool) {
// 	cmd := exec.Command("clear")
// 	cmd.Stdout = os.Stdout
// 	cmd.Run()

// 	fmt.Printf("%d:%s\n", direction.Magnitude, direction.TurnDirection)
// 	drawing := ""
// 	for i, line := range terrain {
// 		for j, char := range line {
// 			if visited[Coordinate{i, j}] {
// 				drawing += "X"
// 			} else {
// 				drawing += char
// 			}
// 		}
// 		drawing += "\n"
// 	}

// 	fmt.Printf("%s", drawing)

// 	input := bufio.NewScanner(os.Stdin)
// 	input.Scan()
// }
