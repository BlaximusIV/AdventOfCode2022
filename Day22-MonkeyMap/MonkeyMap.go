package main

import (
	"fmt"
	"math"
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
	password := getPassword(terrain, directions, false)

	fmt.Printf("2D Password: %d\n", password)

	// Part2
	password = getPassword(terrain, directions, true)

	fmt.Printf("3D Password: %d\n", password)

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

func getPassword(terrain MonkeyMap, directions []Direction, is3D bool) int {
	const RowMult = 1000
	const ColMult = 4

	startCoord := terrain.leftMostCoordinate(0)
	cursor := Cursor{startCoord, Right}

	followDirections(terrain, directions, &cursor, is3D)

	// Column / rows start at 1 for scoring
	rowScore := RowMult * (cursor.Position.Y + 1)
	colScore := ColMult * (cursor.Position.X + 1)
	faceScore := int(cursor.Facing)
	return rowScore + colScore + faceScore
}

func followDirections(terrain MonkeyMap, directions []Direction, cursor *Cursor, is3D bool) {
	for _, direction := range directions {
		// Try to move direction facing
		for i := 0; i < direction.Magnitude; i++ {
			tryMove(terrain, cursor, is3D)
		}

		turn(cursor, direction.TurnDirection)
	}
}

func tryMove(terrain MonkeyMap, cursor *Cursor, is3D bool) {
	coord := cursor.NextCoord()
	plane := terrain.plane(cursor.Position)
	facing := cursor.Facing

	// If invalid, adjust
	switch cursor.Facing {
	case Right:
		{
			if coord.X >= len(terrain[coord.Y]) {
				if is3D {
					coord, facing = crossPlane(plane, cursor.Position, cursor.Facing)
				} else {
					coord = terrain.leftMostCoordinate(coord.Y)
				}
			}
		}
	case Down:
		{
			if coord.Y >= len(terrain) || coord.X >= len(terrain[coord.Y]) {
				if is3D {
					coord, facing = crossPlane(plane, cursor.Position, cursor.Facing)
				} else {
					coord = terrain.topMostCoordinate(coord.X)
				}
			}
		}
	case Left:
		{
			if coord.X < 0 || terrain[coord.Y][coord.X] == " " {
				if is3D {
					coord, facing = crossPlane(plane, cursor.Position, cursor.Facing)
				} else {
					coord = terrain.rightMostCoordinate(coord.Y)
				}
			}
		}
	case Up:
		{
			if coord.Y < 0 || terrain[coord.Y][coord.X] == " " {
				if is3D {
					coord, facing = crossPlane(plane, cursor.Position, cursor.Facing)
				} else {
					coord = terrain.botMostCoordinate(coord.Y, coord.X)
				}
			}
		}
	}

	// If blocked return
	if terrain[coord.Y][coord.X] == "#" {
		return
	}

	// else update cursor Coordinate
	cursor.Position = coord
	cursor.Facing = facing
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

// Only taking into account planes that weren't naturally connected
// Hard coded transitions because I want to move on and no general algorithm comes to mind. Need to revisit to make input-agnostic.
func crossPlane(from int, coord Coordinate, facing Facing) (Coordinate, Facing) {
	switch from {
	case 1:
		{
			// traverse to 6
			if facing == Up {
				return Coordinate{coord.X + 100, coord.Y}, Right
			} else if facing == Left { // 5
				return Coordinate{149 - coord.Y, 0}, Right
			}
		}
	case 2:
		{
			// 6
			if facing == Up {
				return Coordinate{199, coord.X - 100}, Up
			} else if facing == Right { // 4
				return Coordinate{149 - coord.Y, 99}, Left
			} else if facing == Down { // 3
				return Coordinate{coord.X - 50, 99}, Left
			}
		}
	case 3:
		{
			// 2
			if facing == Right {
				return Coordinate{49, coord.Y + 50}, Up
			} else if facing == Left { // 5
				return Coordinate{100, coord.Y - 50}, Down
			}
		}
	case 4:
		{
			// 2
			if facing == Right {
				y := int(math.Abs(float64(149) - float64(coord.Y)))
				return Coordinate{y, 149}, Left
			} else if facing == Down { // 6
				return Coordinate{coord.X + 100, 49}, Left
			}
		}
	case 5:
		{
			// 3
			if facing == Up {
				return Coordinate{coord.X + 50, 50}, Right
			} else if facing == Left { // 1
				return Coordinate{149 - coord.Y, 50}, Right
			}
		}
	case 6:
		{
			// 4
			if facing == Right {
				return Coordinate{149, coord.Y - 100}, Up
			} else if facing == Down { // 2
				return Coordinate{0, coord.X + 100}, Down
			} else if facing == Left { // 1
				return Coordinate{0, coord.Y - 100}, Down
			}
		}
	}

	panic("Trying to traverse to connected pane.")
}
