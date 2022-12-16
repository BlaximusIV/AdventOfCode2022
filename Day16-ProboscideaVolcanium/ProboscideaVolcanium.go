package main

import (
	"log"
	"time"
)

type Node struct {
	Name           string
	FlowRate       int
	ConnectedNodes []Node
}

func main() {
	startTime := time.Now()

	// Go until run out, own copy of visited?

	elapsed := time.Since(startTime)
	log.Printf("Elapsed Time: %s\n", elapsed)
}
