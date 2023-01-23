package main

import "fmt"

type Direction int

const (
	North Direction = 0
	South Direction = 1
	West  Direction = 2
	East  Direction = 3
)

func (d Direction) Next() Direction {
	if d == East {
		return North
	}

	return d + 1
}

type Coordinate struct {
	Y, X int
}

type Elf Coordinate

type Proposal struct {
	From, To Elf
}

type Map map[int]map[int]bool

func (m Map) Add(elf Elf) {
	m.EnsureRow(elf.Y)

	m[elf.Y][elf.X] = true
}

func (m Map) Remove(elf Elf) {
	m.EnsureRow(elf.Y)

	m[elf.Y][elf.X] = false
}

func (m Map) Move(from Elf, to Elf) {
	m.Remove(from)
	m.Add(to)
}

func (m Map) EnsureRow(y int) {
	_, ok := m[y]
	if !ok {
		m[y] = map[int]bool{}
	}
}

func (m Map) HasAnyNeighbors(e Elf) bool {
	return m.HasNeighbors(e, North) || m.HasNeighbors(e, South) || m.HasNeighbors(e, West) || m.HasNeighbors(e, East)
}

func (m Map) HasNeighbors(e Elf, d Direction) bool {
	switch d {
	case North:
		{
			m.EnsureRow(e.Y - 1)
			return m[e.Y-1][e.X-1] || m[e.Y-1][e.X] || m[e.Y-1][e.X+1]
		}
	case South:
		{
			m.EnsureRow(e.Y + 1)
			return m[e.Y+1][e.X-1] || m[e.Y+1][e.X] || m[e.Y+1][e.X+1]
		}
	case West:
		{
			m.EnsureRow(e.Y + 1)
			m.EnsureRow(e.Y - 1)
			return m[e.Y+1][e.X-1] || m[e.Y][e.X-1] || m[e.Y-1][e.X-1]
		}
	case East:
		{
			m.EnsureRow(e.Y + 1)
			m.EnsureRow(e.Y - 1)
			return m[e.Y+1][e.X+1] || m[e.Y][e.X+1] || m[e.Y-1][e.X+1]
		}
	}

	panic("Unknown direction")
}

type Grid struct {
	Map                                      Map
	Elves                                    []Elf
	NorthMost, EastMost, SouthMost, WestMost int
}

func (g *Grid) Move(from, to Elf) {
	g.Map.Move(from, to)
	g.UpdateBounds(to)
}

func (g *Grid) MakeGrid(elves []Elf) {
	m := Map{}
	g.Elves = elves
	g.NorthMost, g.SouthMost, g.EastMost, g.WestMost = elves[0].Y, elves[0].Y, elves[0].X, elves[0].X

	for _, elf := range elves {
		g.UpdateBounds(elf)
		m.Add(elf)
	}

	g.Map = m
}

func (g Grid) CountEmptyTiles() int {
	count := 0
	for i := g.NorthMost; i <= g.SouthMost; i++ {
		for j := g.WestMost; j <= g.EastMost; j++ {
			if !g.Map[i][j] {
				count++
			}
		}
	}

	return count
}

func (g *Grid) UpdateBounds(e Elf) {
	if e.X > g.EastMost {
		g.EastMost = e.X
	}
	if e.Y < g.NorthMost {
		g.NorthMost = e.Y
	}
	if e.X < g.WestMost {
		g.WestMost = e.X
	}
	if e.Y > g.SouthMost {
		g.SouthMost = e.Y
	}
}

func (g Grid) Print() {
	for i := g.NorthMost; i <= g.SouthMost; i++ {
		line := ""
		for j := g.WestMost; j <= g.EastMost; j++ {
			if !g.Map[i][j] {
				line += "."
			} else {
				line += "#"
			}
		}
		line += "\n"
		fmt.Print(line)
	}
	fmt.Print("\n")
}
