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

func (p Point) Euclidian(s Point) int {
	return difference(p.X, s.X) + difference(p.Y, s.Y)
}

func difference(a int, b int) int {
	return int(math.Abs(float64(a - b)))
}

type SensorPair struct {
	Sensor, Beacon Point
	Distance       int
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
	popCount := getDepthMapPopulationCount(sensorList, 2000000)
	log.Printf("%d items intersecting line\n", popCount)

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
	return SensorPair{sensor, beacon, sensor.Euclidian(beacon)}
}

func getDepthMapPopulationCount(sList []SensorPair, targetDepth int) int {
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
		signedDifference := pair.Distance - pair.Sensor.Euclidian(convergencePoint)
		if signedDifference >= 0 {
			r := Range{convergencePoint.X - signedDifference, convergencePoint.X + signedDifference}
			populatedRanges = append(populatedRanges, r)
		}
	}

	// consolidate ranges
	consolidated := getConsolidatedRanges(populatedRanges)

	// Count
	rangeSum := getRangesSum(consolidated)

	// Subtract intersecting points
	rangeSum -= getIntersectingCount(consolidated, intersectingItems)

	return rangeSum
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
