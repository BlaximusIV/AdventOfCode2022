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
	currentVal := g.RegionMap[current.Y][current.X]
	destVal := g.RegionMap[destination.Y][destination.X]
	return g.Elevations[destVal]-g.Elevations[currentVal] <= 1
}

func (g Grid) neighbors(current Coordinate) (neighbors []Coordinate) {
	n := []Coordinate{{current.Y + 1, current.X}, {current.Y - 1, current.X}, {current.Y, current.X - 1}, {current.Y, current.X + 1}}

	for _, val := range n {
		if g.InBounds(val) && g.Passable(current, val) {
			neighbors = append(neighbors, val)
		}
	}

	return
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
