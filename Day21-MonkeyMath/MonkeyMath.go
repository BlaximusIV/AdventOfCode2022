package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Formula struct {
	lhMonkey, rhMonkey, Operand string
}

func main() {
	startTime := time.Now()

	content, _ := os.ReadFile("Input.txt")

	monkeyNumbers := getMonkeyNumbers(string(content))

	// Part 1
	fmt.Printf("Root monkey number: %d\n", monkeyNumbers["root"])

	elapsed := time.Since(startTime)
	fmt.Printf("Elapsed time: %v\n", elapsed)
}

func getMonkeyNumbers(input string) map[string]int {
	monkeyNumbers := map[string]int{}
	monkeyFormulas := map[string]Formula{}
	monkeyLines := strings.Split(input, "\r\n")
	for _, line := range monkeyLines {
		monkeyInstructions := strings.Split(line, " ")
		monkey := strings.Trim(monkeyInstructions[0], ":")

		if len(monkeyInstructions) > 2 {
			monkeyFormulas[monkey] = Formula{monkeyInstructions[1], monkeyInstructions[3], monkeyInstructions[2]}
		} else {
			val, _ := strconv.Atoi(monkeyInstructions[1])
			monkeyNumbers[monkey] = val
		}
	}

	populateMonkey("root", monkeyNumbers, monkeyFormulas)

	return monkeyNumbers
}

func populateMonkey(name string, monkeyNumbers map[string]int, formulas map[string]Formula) {
	if _, ok := monkeyNumbers[name]; ok {
		return
	}

	formula := formulas[name]

	if _, ok := monkeyNumbers[formula.lhMonkey]; !ok {
		populateMonkey(formula.lhMonkey, monkeyNumbers, formulas)
	}

	if _, ok := monkeyNumbers[formula.rhMonkey]; !ok {
		populateMonkey(formula.rhMonkey, monkeyNumbers, formulas)
	}

	monkeyNumbers[name] = getValue(monkeyNumbers[formula.lhMonkey], monkeyNumbers[formula.rhMonkey], formula.Operand)
}

func getValue(lhs, rhs int, operand string) int {
	switch operand {
	case "*":
		{
			return lhs * rhs
		}
	case "+":
		{
			return lhs + rhs
		}
	case "-":
		{
			return lhs - rhs
		}
	case "/":
		{
			return lhs / rhs
		}
	}

	return 0
}
