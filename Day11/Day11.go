/*Solution is not currently clean and was left in the state at which the second part of the puzzle was solved*/
package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Monkey struct {
	Items           []int
	Operation       Operation
	TestMagnitude   int
	PassMonkey      int
	FailMonkey      int
	InspectionCount int
}

type Operation struct {
	Type      string
	Magnitude string
}

func main() {
	start := time.Now()

	content, _ := os.ReadFile("PuzzleInput.txt")
	monkeyLines := strings.Split(string(content), "\r\n")

	monkeys, denominator := populateMonkeys(monkeyLines)

	product := getSimulateMonkeyingProduct(monkeys, 10000, denominator)

	// Part 1
	log.Printf("Product of top two monkey inspector counts: %d", product)

	elapsed := time.Since(start)
	log.Printf("Elapsed Time: %s\n", elapsed)
}

func populateMonkeys(monkeyLines []string) (monkeys []Monkey, commonTestDenominator int) {
	commonTestDenominator = 1
	for i := 0; i < len(monkeyLines); i += 7 {
		monkeyBlock := 7
		if i+monkeyBlock >= len(monkeyLines) {
			monkeyBlock--
		}

		monkey := makeMonkey(monkeyLines[i : i+monkeyBlock])
		commonTestDenominator *= monkey.TestMagnitude
		monkeys = append(monkeys, monkey)
	}

	return
}

func makeMonkey(monkeyLines []string) (monkey Monkey) {
	itemString := strings.Split(strings.Split(monkeyLines[1], ": ")[1], ", ")
	items := []int{}
	for _, item := range itemString {
		value, _ := strconv.Atoi(item)
		items = append(items, value)
	}
	monkey.Items = items

	operationString := strings.Split(monkeyLines[2], " ")
	monkey.Operation = Operation{operationString[len(operationString)-2], operationString[len(operationString)-1]}

	testString := strings.Split(monkeyLines[3], " ")
	testVal, _ := strconv.Atoi(testString[len(testString)-1])
	monkey.TestMagnitude = testVal

	passString := strings.Split(monkeyLines[4], " ")
	passVal, _ := strconv.Atoi(passString[len(passString)-1])
	monkey.PassMonkey = passVal

	failString := strings.Split(monkeyLines[5], " ")
	failVal, _ := strconv.Atoi(failString[len(failString)-1])
	monkey.FailMonkey = failVal

	return
}

func getSimulateMonkeyingProduct(monkeys []Monkey, iterationCount int, commonTestDenominator int) (busyMonkeyProduct int) {
	for i := 0; i < iterationCount; i++ {
		for j, monkey := range monkeys {
			for _, item := range monkeys[j].Items {
				newWorryLevel := item
				var magnitude int
				if monkey.Operation.Magnitude == "old" {
					magnitude = item
				} else {
					parsedMagnitude, _ := strconv.Atoi(monkey.Operation.Magnitude)
					magnitude = parsedMagnitude
				}
				if monkey.Operation.Type == "*" {
					newWorryLevel *= magnitude
				} else {
					newWorryLevel += magnitude
				}

				newWorryLevel = newWorryLevel % commonTestDenominator // int(math.Floor(float64(newWorryLevel) / 3))

				if newWorryLevel%monkey.TestMagnitude == 0 {
					monkeys[monkey.PassMonkey].Items = append(monkeys[monkey.PassMonkey].Items, newWorryLevel)
				} else {
					monkeys[monkey.FailMonkey].Items = append(monkeys[monkey.FailMonkey].Items, newWorryLevel)
				}

				monkeys[j].InspectionCount++
			}

			monkeys[j].Items = []int{}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].InspectionCount > monkeys[j].InspectionCount
	})

	busyMonkeyProduct = monkeys[0].InspectionCount * monkeys[1].InspectionCount

	return
}
