package main

import (
	"math"
	"sort"
)

type Coordinate struct {
	Y, X int
}

func (a Coordinate) EuclidianDistance(b Coordinate) int {
	return difference(a.X, b.X) + difference(a.Y, b.Y)
}

func difference(a int, b int) int {
	return int(math.Abs(float64(a - b)))
}

type State struct {
	Coordinate Coordinate
	Time       int
}

type PriorityItem struct {
	Priority int
	Item     State
}

type PriorityQueue struct {
	Vals []PriorityItem
}

func (q *PriorityQueue) Enqueue(c PriorityItem) {
	q.Vals = append([]PriorityItem{c}, q.Vals...)

	sort.Slice(q.Vals, func(i, j int) bool {
		return q.Vals[i].Priority < q.Vals[j].Priority
	})
}

func (q *PriorityQueue) Dequeue() (c State) {
	c = q.Vals[0].Item
	q.Vals = q.Vals[1:]
	return
}

type Blizzards [][]bool

type Map struct {
	Coordinates              [][]string
	North, South, East, West Blizzards
}

func (m Map) GetMoves(c Coordinate, min int, start Coordinate, end Coordinate) []State {
	moves := []State{}

	time := min + 1
	south := Coordinate{c.Y + 1, c.X}
	if !m.IsOccupied(south, time, start, end) {
		moves = append(moves, State{south, time})
	}

	north := Coordinate{c.Y - 1, c.X}
	if !m.IsOccupied(north, time, start, end) {
		moves = append(moves, State{north, time})
	}

	east := Coordinate{c.Y, c.X + 1}
	if !m.IsOccupied(east, time, start, end) {
		moves = append(moves, State{east, time})
	}

	west := Coordinate{c.Y, c.X - 1}
	if !m.IsOccupied(west, time, start, end) {
		moves = append(moves, State{west, time})
	}

	stay := c
	if !m.IsOccupied(stay, time, start, end) {
		moves = append(moves, State{stay, time})
	}

	return moves
}

func (m Map) IsOccupied(c Coordinate, min int, start Coordinate, end Coordinate) bool {
	if m.IsBorder(c, start, end) {
		return true
	}

	return m.HasBlizzard(c, min)
}

func (m Map) IsBorder(c Coordinate, start Coordinate, end Coordinate) bool {
	if c == end || c == start {
		return false
	}

	return c.Y <= 0 || c.Y >= len(m.Coordinates)-1 || c.X <= 0 || c.X >= len(m.Coordinates[0])-1
}

func (m Map) HasBlizzard(c Coordinate, min int) bool {
	return m.HasBlizzardNorth(c, min) ||
		m.HasBlizzardSouth(c, min) ||
		m.HasBlizzardEast(c, min) ||
		m.HasBlizzardWest(c, min)
}

func (m Map) HasBlizzardNorth(c Coordinate, min int) bool {
	if m.IsVerticalBlizzardSafe(c) {
		return false
	}

	position := (c.Y - 1 + min) % (len(m.Coordinates) - 2)
	return m.North[c.X-1][position]
}

func (m Map) HasBlizzardSouth(c Coordinate, min int) bool {
	if m.IsVerticalBlizzardSafe(c) {
		return false
	}

	blizTravel := min % (len(m.Coordinates) - 2)
	position := (c.Y - 1) - blizTravel
	if position < 0 {
		position = len(m.South[c.X-1]) + position
	}
	return m.South[c.X-1][position]
}

func (m Map) HasBlizzardEast(c Coordinate, min int) bool {
	if m.IsHorizontalBlizzardSafe(c) {
		return false
	}

	blizTravel := min % (len(m.Coordinates[0]) - 2)
	position := (c.X - 1) - blizTravel
	if position < 0 {
		position = len(m.East[c.Y-1]) + position
	}
	return m.East[c.Y-1][position]
}

func (m Map) HasBlizzardWest(c Coordinate, min int) bool {
	if m.IsHorizontalBlizzardSafe(c) {
		return false
	}

	position := (c.X - 1 + min) % (len(m.Coordinates[0]) - 2)
	return m.West[c.Y-1][position]
}

func (m Map) IsHorizontalBlizzardSafe(c Coordinate) bool {
	return c.Y == 0 || c.Y == len(m.Coordinates)-1
}

func (m Map) IsVerticalBlizzardSafe(c Coordinate) bool {
	return c.X == 1 || c.X == len(m.North)
}
