package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Node struct {
	Id, Val        int
	Next, Previous *Node
}

func main() {
	startTime := time.Now()

	input, _ := os.ReadFile("Input.txt")

	numbers := parseInput(string(input))
	ordering, origin := makeNodes(numbers)

	// Part1
	simulateMixing(ordering)

	sum := getGroveCoordinateSum(ordering, origin.Id)
	log.Printf("Grove coordinate sum: %d\n", sum)

	// Part2
	ordering, origin = makeNodes(numbers)

	const DecryptKey = 811589153
	simulateDecryptKeyMixing(ordering, DecryptKey)

	sum = getGroveCoordinateSum(ordering, origin.Id)
	log.Printf("Decrypt key grove coordinate sum: %d\n", sum)

	elapsed := time.Since(startTime)
	log.Printf("Elapsed Time: %s\n", elapsed)
}

func parseInput(input string) []int {
	numbers := []int{}

	for _, str := range strings.Split(input, "\n") {
		i, _ := strconv.Atoi(str)
		numbers = append(numbers, i)
	}

	return numbers
}

// I'm already iterating through it, might as well get everything I need
// Needs refactoring
func makeNodes(numbers []int) (ordering map[int]*Node, origin Node) {
	nodes := []Node{}
	ordering = map[int]*Node{}

	for i, val := range numbers {
		id := i + 1
		node := Node{id, val, nil, nil}
		nodes = append(nodes, node)
	}

	nodes[0].Previous = &nodes[len(nodes)-1]
	nodes[0].Next = &nodes[1]
	nodes[len(nodes)-1].Next = &nodes[0]
	nodes[len(nodes)-1].Previous = &nodes[len(nodes)-2]

	ordering[nodes[0].Id] = &nodes[0]
	ordering[nodes[len(nodes)-1].Id] = &nodes[len(nodes)-1]

	for i := 1; i < len(nodes)-1; i++ {
		nodes[i].Previous = &nodes[i-1]
		nodes[i].Next = &nodes[i+1]
		ordering[nodes[i].Id] = &nodes[i]

		// Depends on the origin node not being the first or last in the list
		if nodes[i].Val == 0 {
			origin = nodes[i]
		}
	}

	return
}

func simulateMixing(ordering map[int]*Node) {
	for i := 1; i < len(ordering)+1; i++ {
		node := ordering[i]
		if node.Val == 0 {
			continue
		}

		// Extract node
		node.Previous.Next = node.Next
		node.Next.Previous = node.Previous

		// Find where it goes
		current := node
		// Modulo to minimize the number of traversals around the entire list
		iterations := node.Val % (len(ordering) - 1)
		for i := iterations; i != 0; increment(&i) {
			current = traverse(current, node.Val)
		}

		// Insert it
		if node.Val > 0 {
			temp := ordering[current.Id].Next
			ordering[current.Id].Next = node
			ordering[temp.Id].Previous = node
			node.Next = ordering[temp.Id]
			node.Previous = ordering[current.Id]
		} else {
			temp := ordering[current.Id].Previous
			ordering[current.Id].Previous = node
			ordering[temp.Id].Next = node
			node.Next = ordering[current.Id]
			node.Previous = ordering[temp.Id]
		}
	}
}

func increment(i *int) {
	if *i > 0 {
		*i--
	} else if *i < 0 {
		*i++
	}
}

func traverse(n *Node, val int) *Node {
	if val > 0 {
		n = n.Next
	} else {
		n = n.Previous
	}

	return n
}

func findPosition(origin *Node, count int) int {
	current := origin
	for i := 0; i < count; i++ {
		current = current.Next
	}

	return current.Val
}

func simulateDecryptKeyMixing(ordering map[int]*Node, key int) {
	for _, node := range ordering {
		node.Val *= key
	}

	const RequiredMixCount = 10
	for i := 0; i < RequiredMixCount; i++ {
		simulateMixing(ordering)
	}
}

func getGroveCoordinateSum(ordering map[int]*Node, id int) int {
	pos1, pos2, pos3 := findPosition(ordering[id], 1000), findPosition(ordering[id], 2000), findPosition(ordering[id], 3000)

	return pos1 + pos2 + pos3
}

// Used for debugging, unused code makes the linter upset
// func reportNodes(origin Node, len int, forwards bool) {
// 	node := &origin
// 	nodes := ""
// 	for i := 0; i < len; i++ {
// 		nodes += strconv.Itoa(node.Val) + " "
// 		if forwards {
// 			node = node.Next
// 		} else {
// 			node = node.Previous
// 		}
// 	}
// 	log.Printf("%s\n", nodes)
// }
