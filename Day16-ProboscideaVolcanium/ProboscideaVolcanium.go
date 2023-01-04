package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()

	content, _ := os.ReadFile("Input.txt")
	graph := getGraph(string(content))

	// Precompile path lengths
	shortestPaths := getShortestPaths(graph)

	// Part 1
	greatestFlow := getSingleGreatestPathsFlow(graph, shortestPaths)
	log.Printf("Greatest possible flow: %v\n", greatestFlow)

	// Part 2
	greatestDuoFlow := getDuoGreatestPathsFlow(graph, shortestPaths)
	log.Printf("Greatest possible duo flow: %v\n", greatestDuoFlow)

	elapsed := time.Since(startTime)
	log.Printf("Elapsed Time: %s\n", elapsed)
}

func getGraph(input string) (graph map[string]Node) {
	graph = map[string]Node{}
	nameExp, _ := regexp.Compile(`([A-Z]{2})`)
	flowExp, _ := regexp.Compile(`(\d)+`)
	for _, line := range strings.Split(input, "\n") {
		nodeNames := nameExp.FindAllString(line, -1)
		flowRateRes := flowExp.FindAllString(line, -1)
		flowRate, _ := strconv.Atoi(flowRateRes[0])

		node := Node{Name: nodeNames[0], FlowRate: flowRate, Edges: nodeNames[1:]}
		graph[node.Name] = node
	}

	return
}

func getShortestPaths(g map[string]Node) (paths map[string]int) {
	paths = map[string]int{}

	for key := range g {
		for key2 := range g {
			if key == key2 {
				continue
			}

			findShortestPath(key, key2, g, paths)
		}
	}

	return
}

func getPathName(origin, destination string) string {
	return origin + " -> " + destination
}

func findShortestPath(origin, destination string, graph map[string]Node, paths map[string]int) {
	// If we've already calculated it from the other direction, set it up and continue
	oppPath, ok := paths[getPathName(destination, origin)]
	pathName := getPathName(origin, destination)
	if ok {
		paths[pathName] = oppPath
		return
	}

	frontier := PriorityQueue{}
	frontier.Enqueue(PriorityItem{0, graph[origin]})
	visited := map[string]string{}
	costSoFar := map[string]int{origin: 0}

	for len(frontier.Vals) > 0 {
		current := frontier.Dequeue()
		if current.Name == destination {
			break
		}

		for _, c := range graph[current.Name].Edges {
			newCost := costSoFar[current.Name] + 1
			_, exists := costSoFar[c]
			if !exists || newCost < costSoFar[c] {
				costSoFar[c] = newCost
				frontier.Enqueue(PriorityItem{newCost, graph[c]})
				visited[c] = current.Name
			}
		}
	}

	path := getRoute(visited, origin, destination)

	paths[pathName] = len(path)
}

func getRoute(visited map[string]string, origin, destination string) (path []string) {
	_, exists := visited[destination]
	if !exists {
		return
	}

	current := destination

	for current != origin {
		path = append(path, current)
		current = visited[current]
	}

	return
}

func getSingleGreatestPathsFlow(graph map[string]Node, paths map[string]int) int {
	releasable := getReleasableValves(graph)

	beginningNode := "AA"
	minutesRemaining := 30
	greatestPathFlow, _ := getGreatestPath(graph, paths, releasable, beginningNode, minutesRemaining, 0)

	return greatestPathFlow
}

func getDuoGreatestPathsFlow(graph map[string]Node, paths map[string]int) int {
	releasable := getReleasableValves(graph)

	// The data is structured in a way where the original path need not vary. This solution does not fit all similar problems.
	beginningNode := "AA"
	minutesRemaining := 26
	greatestPathFlow, greatestPath := getGreatestPath(graph, paths, releasable, beginningNode, minutesRemaining, 0)

	openedNodes := strings.Split(greatestPath, ",")[1:]
	for _, val := range openedNodes {
		releasable[val] = true
	}

	elephlow, _ := getGreatestPath(graph, paths, releasable, "AA", minutesRemaining, 0)

	return greatestPathFlow + elephlow
}

func getGreatestPath(
	graph map[string]Node,
	paths map[string]int,
	targets map[string]bool,
	current string,
	ttl,
	score int) (int, string) {

	if ttl <= 0 {
		return score, ""
	}

	maxScore := 0
	maxPath := ""
	for key, val := range targets {
		path := getPathName(current, key)
		// time to travel, along with time to activate
		activationCost := paths[path] + 1
		if !val && activationCost < ttl {
			copy := copyMap(targets)
			copy[key] = true
			newTtl := ttl - activationCost
			newScore := newTtl * graph[key].FlowRate
			resultScore, resultPath := getGreatestPath(graph, paths, copy, key, newTtl, newScore)

			if resultScore > maxScore {
				maxScore = resultScore
				maxPath = resultPath
			}
		}
	}

	return score + maxScore, current + "," + maxPath
}

func getReleasableValves(graph map[string]Node) map[string]bool {
	releasable := map[string]bool{}
	for key, val := range graph {
		if val.FlowRate > 0 {
			releasable[key] = false
		}
	}

	return releasable
}

func copyMap(m map[string]bool) (copy map[string]bool) {
	copy = map[string]bool{}

	for key, val := range m {
		copy[key] = val
	}

	return
}
