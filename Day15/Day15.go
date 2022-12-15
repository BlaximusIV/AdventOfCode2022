/* Desperately needs refactoring. Can be made more efficient, as well as clean. */
package main

import (
	"log"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	X, Y int
}

func (p Point) EuclidianDistance(s Point) int {
	return difference(p.X, s.X) + difference(p.Y, s.Y)
}

func difference(a int, b int) int {
	return int(math.Abs(float64(a - b)))
}

type SensorPair struct {
	Sensor, Beacon    Point
	EuclidianDistance int
}

type Range struct {
	X1, X2 int
}

func (r Range) Intersects(r2 Range) bool {
	return r.IsInRange(r2.X1) || r.IsInRange(r2.X2) || r2.IsInRange(r.X1) || r2.IsInRange(r.X2)
}

func (r Range) IsInRange(value int) bool {
	return value >= r.X1 && value <= r.X2
}

func main() {
	startTime := time.Now()

	content, _ := os.ReadFile("Input.txt")
	sensorList := []SensorPair{}
	populateSensors(&sensorList, string(content))

	// Part1
	popCount, _ := getDepthMapPopulationCount(sensorList, 2000000)
	log.Printf("%d items intersecting line\n", popCount)

	// Part2
	frequency := findMissingBeaconFrequency(sensorList)
	log.Printf("Missing beacon frequency: %d\n", frequency)

	elapsed := time.Since(startTime)
	log.Printf("Elapsed Time: %s\n", elapsed)
}

func populateSensors(sList *[]SensorPair, input string) {
	r, _ := regexp.Compile(`([xy]=-?[\d]*)`)
	for _, line := range strings.Split(input, "\r\n") {
		coordinates := r.FindAllString(line, -1)
		pair := extractPair(coordinates)
		*sList = append(*sList, pair)
	}
}

// In this case, I know the exact length
func extractPair(rawVals []string) SensorPair {
	vals := []int{}
	for _, rawVal := range rawVals {
		val, _ := strconv.Atoi(strings.Split(rawVal, "=")[1])
		vals = append(vals, val)
	}
	sensor, beacon := Point{vals[0], vals[1]}, Point{vals[2], vals[3]}
	return SensorPair{sensor, beacon, sensor.EuclidianDistance(beacon)}
}

func findMissingBeaconFrequency(sensors []SensorPair) int {
	// For 0..4000000
	rangeMin := 0
	rangeMax := 4000000
	targetPoint := Point{0, 0}
	for i := rangeMin; i <= rangeMax; i++ {
		// Get line
		_, ranges := getDepthMapPopulationCount(sensors, i)

		// Trim it
		trimmedRanges := getTrimmedRanges(ranges, rangeMin, rangeMax)

		// if there's more than one range, get the hole
		if len(trimmedRanges) > 1 {
			targetPoint.X = ranges[0].X2 + 1
			targetPoint.Y = i
			break
		}

	}

	// return calculated frequency
	return (4000000 * targetPoint.X) + targetPoint.Y
}

func getDepthMapPopulationCount(sList []SensorPair, targetDepth int) (rangeSum int, ranges []Range) {
	populatedRanges := []Range{}
	intersectingItems := []int{}

	for _, pair := range sList {
		if pair.Beacon.Y == targetDepth && !contains(intersectingItems, pair.Beacon.X) {
			intersectingItems = append(intersectingItems, pair.Beacon.X)
		} else if pair.Sensor.Y == targetDepth && !contains(intersectingItems, pair.Sensor.X) {
			intersectingItems = append(intersectingItems, pair.Sensor.X)
		}

		// if the point is even in range
		// get range
		convergencePoint := Point{pair.Sensor.X, targetDepth}
		signedDifference := pair.EuclidianDistance - pair.Sensor.EuclidianDistance(convergencePoint)
		if signedDifference >= 0 {
			r := Range{convergencePoint.X - signedDifference, convergencePoint.X + signedDifference}
			populatedRanges = append(populatedRanges, r)
		}
	}

	// consolidate ranges
	ranges = getConsolidatedRanges(populatedRanges)

	// Count
	rangeSum = getRangesSum(ranges)

	// Subtract intersecting points
	rangeSum -= getIntersectingCount(ranges, intersectingItems)

	return
}

func getTrimmedRanges(ranges []Range, min int, max int) (trimmedRange []Range) {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].X1 < ranges[j].X1
	})

	minMax := Range{min, max}
	for _, r := range ranges {
		if r.X1 >= min && r.X2 <= max {
			trimmedRange = append(trimmedRange, r)
			continue
		}
		if r.X1 <= min && r.X2 >= max {
			trimmedRange = append(trimmedRange, Range{min, max})
			break
		}
		if r.X1 <= min && minMax.IsInRange(r.X2) {
			trimmedRange = append(trimmedRange, Range{min, r.X2})
			continue
		}
		//
		if minMax.IsInRange(r.X1) && r.X2 >= max {
			trimmedRange = append(trimmedRange, Range{r.X1, max})
			break
		}
	}

	return
}

func contains(items []int, item int) bool {
	for _, val := range items {
		if val == item {
			return true
		}
	}
	return false
}

func getConsolidatedRanges(original []Range) (consolidated []Range) {
	sort.Slice(original, func(i, j int) bool {
		return original[i].X1 < original[j].X1
	})

	temp := original[0]
	for i := 1; i < len(original); i++ {
		if temp.Intersects(original[i]) {
			temp = getCombinedRange(temp, original[i])
		} else {
			consolidated = append(consolidated, temp)
			temp = original[i]
		}
	}

	consolidated = append(consolidated, temp)

	return
}

func getCombinedRange(r1 Range, r2 Range) Range {
	combined := Range{r1.X1, r1.X2}
	if r2.X1 < combined.X1 {
		combined.X1 = r2.X1
	}

	if r2.X2 > combined.X2 {
		combined.X2 = r2.X2
	}

	return combined
}

func getRangesSum(r []Range) (sum int) {
	sum = 0

	for _, val := range r {
		sum += (val.X2 - val.X1)
		// The range technically contains one more space if it crosses into negative
		if val.X2 > 0 && val.X1 < 0 {
			sum++
		}
	}

	return
}

func getIntersectingCount(r []Range, xVals []int) (sum int) {
	for _, val := range xVals {
		for _, item := range r {
			if item.IsInRange(val) {
				sum++
				continue
			}
		}
	}

	return
}
