package main

import "sort"

type Coordinate struct {
	Y, X int
}

type Grid struct {
	Height, Width int
	RegionMap     [][]string
	Elevations    map[string]int
}

func (g Grid) InBounds(coord Coordinate) bool {
	return 0 <= coord.X && coord.X < g.Width && 0 <= coord.Y && coord.Y < g.Height
}

func (g Grid) Passable(current Coordinate, destination Coordinate) bool {
	return g.GetElevation(destination)-g.GetElevation(current) <= 1
}

func (g Grid) Neighbors(current Coordinate) (neighbors []Coordinate) {
	n := []Coordinate{{current.Y + 1, current.X}, {current.Y - 1, current.X}, {current.Y, current.X - 1}, {current.Y, current.X + 1}}

	for _, val := range n {
		if g.InBounds(val) && g.Passable(current, val) {
			neighbors = append(neighbors, val)
		}
	}

	return
}

func (g Grid) GetElevation(c Coordinate) int {
	return g.Elevations[g.RegionMap[c.Y][c.X]]
}

type PriorityItem struct {
	Priority int
	Item     Coordinate
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

func (q *PriorityQueue) Dequeue() (c Coordinate) {
	c = q.Vals[0].Item
	q.Vals = q.Vals[1:]
	return
}
