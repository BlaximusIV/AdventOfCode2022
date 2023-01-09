// Types created for the problem along with functions for those types

package main

type Block struct {
	Points []Point
}

type State struct {
	BlockTop                 [][7]string
	WindIndex, NextIteration int
}

func (s State) Equals(s2 State) bool {
	if s.WindIndex != s2.WindIndex {
		return false
	}

	if s.NextIteration%5 != s2.NextIteration%5 {
		return false
	}

	for i := 0; i < len(s.BlockTop); i++ {
		for j := 0; j < len(s.BlockTop[i]); j++ {
			if s.BlockTop[i][j] != s2.BlockTop[i][j] {
				return false
			}
		}
	}

	return true
}

func (b Block) FindLeft() int {
	left := b.Points[0].X
	for i := 1; i < len(b.Points); i++ {
		if b.Points[i].X < left {
			left = b.Points[i].X
		}
	}
	return left
}

func (b Block) FindRight() int {
	right := b.Points[0].X
	for i := 1; i < len(b.Points); i++ {
		if b.Points[i].X > right {
			right = b.Points[i].X
		}
	}
	return right
}

func (b Block) FindBottom() int {
	bottom := b.Points[0].Y
	for i := 1; i < len(b.Points); i++ {
		if b.Points[i].Y > bottom {
			bottom = b.Points[i].Y
		}
	}
	return bottom
}

type Point struct {
	Y, X int
}

func findHighestBlock(chamber [][7]string) int {
	for i, level := range chamber {
		for _, val := range level {
			if val != SpaceDelimeter {
				return i
			}
		}
	}

	// Is empty
	return len(chamber)
}
