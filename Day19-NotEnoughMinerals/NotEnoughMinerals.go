package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type BluePrint struct {
	Id                                               int
	OrebotCost, ClaybotCost, ObsibotCost, GeobotCost [3]int
}

func main() {
	startTime := time.Now()

	content, _ := os.ReadFile("Input.txt")

	bluePrints := parseBluePrints(string(content))

	// Part1
	qualityLevels := getQualityLevels(bluePrints)
	sum := 0
	for _, level := range qualityLevels {
		sum += level
	}

	log.Printf("Quality level sum: %d\n", sum)

	elapsed := time.Since(startTime)
	log.Printf("Elapsed Time: %s\n", elapsed)
}

func parseBluePrints(input string) []BluePrint {
	bluePrints := []BluePrint{}

	valueExp, _ := regexp.Compile(`(\d)+`)
	for _, line := range strings.Split(input, "\n") {
		intVals := []int{}
		values := valueExp.FindAllString(line, -1)
		for _, val := range values {
			value, _ := strconv.Atoi(val)

			intVals = append(intVals, value)
		}

		oreBot, clayBot, obsiBot, geoBot := [3]int{}, [3]int{}, [3]int{}, [3]int{}
		oreBot[0], clayBot[0], obsiBot[0], obsiBot[1], geoBot[0], geoBot[2] = intVals[1], intVals[2], intVals[3], intVals[4], intVals[5], intVals[6]

		bluePrints = append(bluePrints, BluePrint{intVals[0], oreBot, clayBot, obsiBot, geoBot})
	}

	return bluePrints
}

func getQualityLevels(bluePrints []BluePrint) (levels []int) {
	levels = []int{}

	for _, bp := range bluePrints {
		maxGeodes := getMaxGeodes(bp, 24, [4]int{1, 0, 0, 0}, [4]int{})
		levels = append(levels, bp.Id*maxGeodes)
	}

	return
}

func getMaxGeodes(bp BluePrint, ttl int, bots, resources [4]int) int {
	if ttl <= 0 {
		return resources[3]
	}

	maxGeodes := 0
	choices := getBotChoices(bots)
	for botType, available := range choices {
		if !available {
			continue
		}

		botCopy := copy(bots)
		resourceCopy := copy(resources)
		ttlCopy := ttl

		purchaseBot(bp, &botCopy, &resourceCopy, botType, &ttlCopy)

		botMax := getMaxGeodes(bp, ttlCopy, botCopy, resourceCopy)
		if botMax > maxGeodes {
			maxGeodes = botMax
		}
	}

	return maxGeodes
}

func getBotChoices(currentBots [4]int) [4]bool {
	choices := [4]bool{true, true, false, false}

	// We have a clay robot
	if currentBots[1] > 0 {
		choices[2] = true
	}

	// we have an obsidian bot
	if currentBots[2] > 0 {
		choices[3] = true
	}

	return choices
}

// Used because just using a slice doesn't have the array size guarantee
func copy(a [4]int) [4]int {
	copy := [4]int{a[0], a[1], a[2], a[3]}
	return copy
}

func purchaseBot(bp BluePrint, bots, res *[4]int, botType int, ttl *int) {
	// bots & resources: ore | clay | obsidian | geode
	switch botType {
	case 0:
		{
			// TODO: calculate time, resources necessary
			for res[0] < bp.OrebotCost[0] && *ttl > 1 {
				addResources(bots, res, ttl)
			}

			res[0] -= bp.OrebotCost[0]
			addResources(bots, res, ttl)
			bots[0]++
		}
	case 1:
		{
			for res[0] < bp.ClaybotCost[0] && *ttl > 1 {
				addResources(bots, res, ttl)
			}
			res[0] -= bp.ClaybotCost[0]
			addResources(bots, res, ttl)
			bots[1]++
		}
	case 2:
		{
			for (res[0] < bp.ObsibotCost[0] || res[1] < bp.ObsibotCost[1]) && *ttl > 1 {
				addResources(bots, res, ttl)
			}
			res[0] -= bp.ObsibotCost[0]
			res[1] -= bp.ObsibotCost[1]
			addResources(bots, res, ttl)
			bots[2]++
		}
	case 3:
		{
			for (res[0] < bp.GeobotCost[0] || res[2] < bp.GeobotCost[2]) && *ttl > 1 {
				addResources(bots, res, ttl)
			}
			res[0] -= bp.GeobotCost[0]
			res[2] -= bp.GeobotCost[2]
			addResources(bots, res, ttl)
			bots[3]++
		}
	}
}

// add by time?
func addResources(bots, res *[4]int, ttl *int) {
	for i, val := range bots {
		res[i] += val
	}
	*ttl--
}

// func max(a, b int) int {
// 	max := int(math.Max(float64(a), float64(b)))
// 	return max
// }
