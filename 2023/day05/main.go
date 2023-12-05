package main

import (
	"fmt"
	"sort"

	"github.com/nikolovlazar/aoc-go/input"
	"github.com/nikolovlazar/aoc-go/utils"
)

type DestinationLength struct {
	destination, length int
}

func main() {
	// Read the input
	lines := input.ReadInput("/2023/day05/input_small.txt")

	// part1(lines)
	part2(lines)
}

func part1(lines []string) {
	seeds := []int{}

	seedToSoil := map[int]DestinationLength{}
	soilToFertilizer := map[int]DestinationLength{}
	fertilizerToWater := map[int]DestinationLength{}
	waterToLight := map[int]DestinationLength{}
	lightToTemperature := map[int]DestinationLength{}
	temperatureToHumidity := map[int]DestinationLength{}
	humidityToLocation := map[int]DestinationLength{}

	inputFlag := ""
	inputGroups := []string{
		"seed-to-soil map:",
		"soil-to-fertilizer map:",
		"fertilizer-to-water map:",
		"water-to-light map:",
		"light-to-temperature map:",
		"temperature-to-humidity map:",
		"humidity-to-location map:",
	}

	for index, line := range lines {
		if index == 0 {
			seedNumbersString := line[7:]
			seeds = utils.ExtractNumbers(seedNumbersString)
			continue
		}

		if line == "" {
			continue
		}

		if utils.Contains(inputGroups, line) {
			inputFlag = line
			continue
		}

		numbers := utils.ExtractNumbers(line)
		destinationStart, sourceStart, length := numbers[0], numbers[1], numbers[2]

		newMap := createMap(destinationStart, sourceStart, length)

		switch inputFlag {
		case "seed-to-soil map:":
			seedToSoil = utils.MergeMaps(seedToSoil, newMap)
		case "soil-to-fertilizer map:":
			soilToFertilizer = utils.MergeMaps(soilToFertilizer, newMap)
		case "fertilizer-to-water map:":
			fertilizerToWater = utils.MergeMaps(fertilizerToWater, newMap)
		case "water-to-light map:":
			waterToLight = utils.MergeMaps(waterToLight, newMap)
		case "light-to-temperature map:":
			lightToTemperature = utils.MergeMaps(lightToTemperature, newMap)
		case "temperature-to-humidity map:":
			temperatureToHumidity = utils.MergeMaps(temperatureToHumidity, newMap)
		case "humidity-to-location map:":
			humidityToLocation = utils.MergeMaps(humidityToLocation, newMap)
		}
	}

	smallestLocation := -1
	for _, seedNumber := range seeds {
		soilNumber := matchNumber(seedToSoil, seedNumber)
		fertilizerNumber := matchNumber(soilToFertilizer, soilNumber)
		waterNumber := matchNumber(fertilizerToWater, fertilizerNumber)
		lightNumber := matchNumber(waterToLight, waterNumber)
		temperatureNumber := matchNumber(lightToTemperature, lightNumber)
		humidityNumber := matchNumber(temperatureToHumidity, temperatureNumber)
		locationNumber := matchNumber(humidityToLocation, humidityNumber)

		if smallestLocation == -1 {
			smallestLocation = locationNumber
			continue
		}
		if locationNumber < smallestLocation {
			smallestLocation = locationNumber
		}
	}

	fmt.Printf("Part 1: %d", smallestLocation)
}

func part2(lines []string) {
	seeds := map[int]int{}

	seedToSoil := [][]int{}
	soilToFertilizer := [][]int{}
	fertilizerToWater := [][]int{}
	waterToLight := [][]int{}
	lightToTemperature := [][]int{}
	temperatureToHumidity := [][]int{}
	humidityToLocation := [][]int{}

	inputFlag := ""
	inputGroups := []string{
		"seed-to-soil map:",
		"soil-to-fertilizer map:",
		"fertilizer-to-water map:",
		"water-to-light map:",
		"light-to-temperature map:",
		"temperature-to-humidity map:",
		"humidity-to-location map:",
	}

	for index, line := range lines {
		if index == 0 {
			seedNumbersString := line[7:]
			seeds = parseSeedNumbers(seedNumbersString)
			continue
		}

		if line == "" {
			continue
		}

		if utils.Contains(inputGroups, line) {
			inputFlag = line
			continue
		}

		numbers := utils.ExtractNumbers(line)

		switch inputFlag {
		case "seed-to-soil map:":
			seedToSoil = append(seedToSoil, numbers)
		case "soil-to-fertilizer map:":
			soilToFertilizer = append(soilToFertilizer, numbers)
		case "fertilizer-to-water map:":
			fertilizerToWater = append(fertilizerToWater, numbers)
		case "water-to-light map:":
			waterToLight = append(waterToLight, numbers)
		case "light-to-temperature map:":
			lightToTemperature = append(lightToTemperature, numbers)
		case "temperature-to-humidity map:":
			temperatureToHumidity = append(temperatureToHumidity, numbers)
		case "humidity-to-location map:":
			humidityToLocation = append(humidityToLocation, numbers)
		}
	}

	smallestLocation := -1
	categories := [][][]int{
		seedToSoil,
		soilToFertilizer,
		fertilizerToWater,
		waterToLight,
		lightToTemperature,
		temperatureToHumidity,
		humidityToLocation,
	}

	for start, length := range seeds {
		rng := []int{start, length}
		locationRanges := convertThroughCategories(rng, categories)

		for _, locationRange := range locationRanges {
			locationStart := locationRange[0]
			if smallestLocation > locationStart || smallestLocation == -1 {
				smallestLocation = locationStart
			}
		}
	}

	fmt.Printf("Part 2: %d", smallestLocation)
}

func convertThroughCategories(seedRange []int, categories [][][]int) [][]int {
	currentRanges := [][]int{seedRange}

	for _, category := range categories {
		newRanges := [][]int{}

		for _, currentRange := range currentRanges {
			convertedRanges := convertRange(currentRange, category)
			newRanges = append(newRanges, convertedRanges...)
		}

		currentRanges = newRanges
	}

	return currentRanges
}

func convertRange(seedRange []int, numbersMap [][]int) [][]int {
	start, length := seedRange[0], seedRange[1]

	ranges := [][]int{}

	for _, row := range numbersMap {
		destinationStart, srcStart, rangeLength := row[0], row[1], row[2]
		srcEnd := srcStart + rangeLength

		if start < srcEnd && start+length > srcStart {
			// We have an overlap
			overlapStart := max(start, srcStart)
			overlapEnd := min(start+length, srcEnd)
			newStart := destinationStart + (overlapStart - srcStart)

			newRange := []int{newStart, overlapEnd - overlapStart}
			ranges = append(ranges, newRange)
		}
	}

	if len(ranges) > 0 {
		return ranges
	} else {
		return append([][]int{}, seedRange)
	}
}

func parseSeedNumbers(seedNumbersString string) map[int]int {
	seedNumbers := utils.ExtractNumbers(seedNumbersString)
	numbers := map[int]int{}

	for i := 0; i < len(seedNumbers); i += 2 {
		start := seedNumbers[i]
		length := seedNumbers[i+1]

		numbers[start] = length
	}

	return numbers
}

func matchNumber(numberMap map[int]DestinationLength, number int) int {
	value, ok := numberMap[number]

	if ok {
		return value.destination
	}

	ranges := []int{}
	for key := range numberMap {
		ranges = append(ranges, key)
	}

	potentialRangeStart := findPotentialRangeStart(ranges, number)

	if potentialRangeStart > number {
		return number
	}

	matchedNumber := number

	destination := numberMap[potentialRangeStart].destination
	length := numberMap[potentialRangeStart].length
	if potentialRangeStart+length > number {
		difference := number - potentialRangeStart
		matchedNumber = destination + difference
	}

	return matchedNumber
}

func findPotentialRangeStart(ranges []int, number int) int {
	sort.Ints(ranges)

	potentialStart := ranges[0]

	for _, start := range ranges {
		if start > potentialStart && start <= number {
			potentialStart = start
		}
	}

	return potentialStart
}

func createMap(destination, source, length int) map[int]DestinationLength {
	value := DestinationLength{destination, length}
	result := map[int]DestinationLength{}

	result[source] = value

	return result
}
