package main

type MonkeyMap [][]string

// Might be able to condense these into a function that takes functions
// I made the call that I don't need to worry about getting out of bounds, I have the data I will run against
func (m MonkeyMap) leftMostCoordinate(y int) Coordinate {
	x := 0
	for m[y][x] == " " {
		x++
	}
	return Coordinate{y, x}
}

func (m MonkeyMap) rightMostCoordinate(y int) Coordinate {
	x := len(m[y]) - 1
	return Coordinate{y, x}
}

func (m MonkeyMap) topMostCoordinate(x int) Coordinate {
	y := 0
	for m[y][x] == " " {
		y++
	}
	return Coordinate{y, x}
}

func (m MonkeyMap) botMostCoordinate(y, x int) Coordinate {
	found := false
	for !found {
		if y == len(m)-1 {
			// Bottom row
			found = true
			continue
		} else if len(m[y+1])-1 < x {
			// Next row doesn't have an x of that value
			found = true
			continue
		}
		y++
	}
	return Coordinate{y, x}
}

type Direction struct {
	Magnitude     int
	TurnDirection string
}

type Coordinate struct {
	Y, X int
}

type Facing int

const (
	Right Facing = 0
	Down  Facing = 1
	Left  Facing = 2
	Up    Facing = 3
)

type Cursor struct {
	Position Coordinate
	Facing   Facing
}

func (c Cursor) NextCoord() Coordinate {
	switch c.Facing {
	case Right:
		return Coordinate{c.Position.Y, c.Position.X + 1}
	case Down:
		return Coordinate{c.Position.Y + 1, c.Position.X}
	case Left:
		return Coordinate{c.Position.Y, c.Position.X - 1}
	case Up:
		return Coordinate{c.Position.Y - 1, c.Position.X}
	}

	return Coordinate{}
}
