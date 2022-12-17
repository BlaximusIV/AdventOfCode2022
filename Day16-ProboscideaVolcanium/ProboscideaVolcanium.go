package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

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

func main() {
	startTime := time.Now()

	content, _ := os.ReadFile("TestInput.txt")
	graph := getGraph(string(content))

	// Part 1
	greatestFlow := getGreatestPathsFlow(graph["AA"], graph, []string{}, map[string]bool{}, 30, 0, 0)
	log.Printf("Greatest possible flow: %v\n", greatestFlow)

	elapsed := time.Since(startTime)
	log.Printf("Elapsed Time: %s\n", elapsed)
}

func getGraph(input string) (graph map[string]Node) {
	graph = map[string]Node{}
	nameExp, _ := regexp.Compile(`([A-Z]{2})`)
	flowExp, _ := regexp.Compile(`(\d)+`)
	for _, line := range strings.Split(input, "\r\n") {
		nodeNames := nameExp.FindAllString(line, -1)
		flowRateRes := flowExp.FindAllString(line, -1)
		flowRate, _ := strconv.Atoi(flowRateRes[0])

		node := Node{Name: nodeNames[0], FlowRate: flowRate, Edges: nodeNames[1:]}
		graph[node.Name] = node
	}

	return
}

func getGreatestPathsFlow(current Node,
	graph map[string]Node,
	path []string,
	activated map[string]bool,
	ttl int,
	currentFlowRate int,
	currentFlowTotal int) int {

	currentFlowTotal += currentFlowRate
	if ttl <= 0 {
		return currentFlowTotal
	}

	path = append(path, current.Name)

	_, exists := activated[current.Name]
	if !exists {
		activated[current.Name] = false
	}

	greatestFlow := 0
	canActivate := current.FlowRate > 0 && ttl > 1 && !activated[current.Name]
	if canActivate {
		newActivate := activated
		newActivate[current.Name] = true
		for _, node := range current.Edges {
			newRate := currentFlowRate + graph[node].FlowRate
			pathFlow := getGreatestPathsFlow(graph[node], graph, path, newActivate, ttl-2, newRate, currentFlowTotal+currentFlowRate)
			if greatestFlow < pathFlow {
				greatestFlow = pathFlow
			}
		}
	}
	for _, node := range current.Edges {
		pathFlow := getGreatestPathsFlow(graph[node], graph, path, activated, ttl-1, currentFlowRate, currentFlowTotal)
		if greatestFlow < pathFlow {
			greatestFlow = pathFlow
		}
	}

	return greatestFlow
}

// func getPaths(current Node, p Path, graph map[string]Node) []Path {
// 	// Update visited
// 	_, exists := p.Visited[current.Name]
// 	if !exists {
// 		p.Visited[current.Name] = false
// 	}

// 	// If no more ttl return path array with single path
// 	paths := []Path{}
// 	if p.TTL == 0 {
// 		paths = append(paths, p)
// 		return paths
// 	}

// 	canActivate := graph[current.Name].FlowRate > 0 && p.TTL > 1 && !p.Visited[current.Name]
// 	for _, node := range graph[current.Name].Edges {
// 		if canActivate {
// 			activatePath := p
// 			activatePath.Visited[current.Name] = true
// 			activatePath.TTL -= 2
// 			activatePath.TotalFlow = activatePath.TTL * graph[current.Name].FlowRate

// 			paths = append(paths, getPaths(graph[node], activatePath, graph)...)
// 		}

// 		p.TTL--
// 		paths = append(paths, getPaths(graph[node], p, graph)...)
// 	}

// 	return paths
// }
