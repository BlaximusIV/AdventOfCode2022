package main

import "sort"

type Node struct {
	Name     string
	FlowRate int
	Edges    []string
}

type Path struct {
	Visited   []string
	TTL       int
	TotalFlow int
}

type PriorityItem struct {
	Priority int
	Item     Node
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

func (q *PriorityQueue) Dequeue() (c Node) {
	c = q.Vals[0].Item
	q.Vals = q.Vals[1:]
	return
}
