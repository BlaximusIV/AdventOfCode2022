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

	numbers, formulas := getMonkeyInfo(string(content))

	// Part 1
	pOneNumbers := copyNumbers(numbers)
	populateMonkey("root", pOneNumbers, formulas)
	fmt.Printf("Root monkey number: %d\n", pOneNumbers["root"])

	// Part 2
	equalizingNumber := getEqualizingNumber(numbers, formulas)
	fmt.Printf("Equalizing number: %d\n", equalizingNumber)

	elapsed := time.Since(startTime)
	fmt.Printf("Elapsed time: %v\n", elapsed)
}

func getMonkeyInfo(input string) (monkeyNumbers map[string]int, monkeyFormulas map[string]Formula) {
	monkeyNumbers = map[string]int{}
	monkeyFormulas = map[string]Formula{}
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

	return
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

func getEqualizingNumber(monkeyNumbers map[string]int, monkeyFormulas map[string]Formula) int {
	guess := 1
	leftMonkey := 1
	rightMonkey := 0
	// right monkey is not dependent on the humn value
	// Find the point at which lh is less, narrow from there (Currently only functional for real input)
	for {
		leftMonkey, rightMonkey = getGuessResult(monkeyNumbers, monkeyFormulas, guess)

		if leftMonkey < rightMonkey {
			break
		}

		guess *= 2
	}

	// Now we know that the correct number is between half the guess number and the current guess
	low := guess / 2
	high := guess
	for leftMonkey != rightMonkey {
		mid := low + ((high - low) / 2)

		midResult, _ := getGuessResult(monkeyNumbers, monkeyFormulas, mid)

		// We know it's too high
		if midResult < rightMonkey {
			high = mid
		} else if midResult > rightMonkey {
			low = mid
		} else {
			guess = mid
		}
		leftMonkey = midResult

		fmt.Printf("Difference: %d\n", leftMonkey-rightMonkey)
	}

	return guess
}

func getGuessResult(monkeyNumbers map[string]int, monkeyFormulas map[string]Formula, guess int) (lh int, rh int) {
	rootFormula := monkeyFormulas["root"]
	copy := copyNumbers(monkeyNumbers)
	copy["humn"] = guess

	populateMonkey(rootFormula.lhMonkey, copy, monkeyFormulas)
	populateMonkey(rootFormula.rhMonkey, copy, monkeyFormulas)

	lh = copy[rootFormula.lhMonkey]
	rh = copy[rootFormula.rhMonkey]

	return
}

// There should be a better way to do a deep copy in go. I don't know what it is yet. Need to move on.
func copyNumbers(monkeyNumbers map[string]int) (copy map[string]int) {
	copy = map[string]int{}

	for key, val := range monkeyNumbers {
		copy[key] = val
	}
	return
}
