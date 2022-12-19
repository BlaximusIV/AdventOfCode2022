package main

import (
	"log"
	"os"
	"time"
)

func main() {
	startTime := time.Now()

	content, _ := os.ReadFile("TestInput.txt")

	elapsed := time.Since(startTime)
	log.Printf("Elapsed Time: %s\n", elapsed)
}
