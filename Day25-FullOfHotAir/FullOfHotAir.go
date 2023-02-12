package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()

	input, _ := os.ReadFile("Input.txt")

	snafuNumbers := strings.Split(string(input), "\n")

	// Part 1
	sum := sumSnafuNumbers(snafuNumbers)
	fmt.Printf("Snafu number sum: %d\n", sum)
	snafu := DecimalToSnafu(sum)
	fmt.Printf("Sum as snafu: %s\n", snafu)

	elapsed := time.Since(startTime)
	fmt.Printf("Elapsed time: %v\n", elapsed)
}

func sumSnafuNumbers(snafuNumbers []string) int {
	sum := 0

	for _, val := range snafuNumbers {
		sum += SnafuToDecimal(val)
	}

	return sum
}

func SnafuToDecimal(snafu string) int {
	sum := 0

	powers := len(snafu) - 1
	for _, val := range strings.Split(snafu, "") {
		value := GetDecimalRepresentation(val)
		if powers == 0 {
			sum += value
		} else {
			sum += int(math.Pow(5, float64(powers)) * float64(value))
		}
		powers--
	}

	return sum
}

func GetDecimalRepresentation(character string) int {
	switch character {
	case "-":
		{
			return -1
		}
	case "=":
		{
			return -2
		}
	case "1":
		{
			return 1
		}
	case "2":
		{
			return 2
		}
	case "0":
		{
			return 0
		}
	}
	panic("Unknown character supplied")
}

func DecimalToSnafu(number int) string {
	snafu := ""

	power := 0
	for SnafuMaxNumber(int(power)) < number {
		power++
	}

	baseNumber := int(math.Pow(5, float64(power)))
	if baseNumber+SnafuMaxNumber(int(power-1)) < number {
		baseNumber *= 2
		snafu += "2"
	} else {
		snafu += "1"
	}
	// Find the rest
	remainder := number - baseNumber
	for remainder != 0 && power >= 0 {
		power--
		val := SnafuPow(power)
		lowerVal := SnafuMaxNumber(power - 1)
		if lowerVal > Abs(remainder) {
			snafu += "0"
		} else if remainder < 0 {
			if val+lowerVal >= Abs(remainder) {
				snafu += "-"
			} else {
				val *= 2
				snafu += "="
			}
			remainder += val
		} else if remainder > 0 {
			if val+lowerVal >= Abs(remainder) {
				snafu += "1"
			} else {
				val *= 2
				snafu += "2"
			}
			remainder -= val
		}
	}

	for power > 0 {
		snafu += "0"
		power--
	}

	return snafu
}

func SnafuMaxNumber(power int) int {
	sum := 0
	for power >= 0 {
		sum += 2 * SnafuPow(power)
		power--
	}

	return sum
}

func Abs(value int) int {
	return int(math.Abs(float64(value)))
}

func SnafuPow(power int) int {
	return (int(math.Pow(5, float64(power))))
}
