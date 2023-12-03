package main

import (
	"fmt"
	"strconv"

	"github.com/nikolovlazar/aoc-go/input"
)

type Coord struct {
	X, Y int
}

func main() {
	// Read the input
	lines := input.ReadInput("/2023/day03/input.txt")

	part1(lines)
	part2(lines)
}

func part2(lines []string) {
	numbersMap := generateNumbersMap(lines)

	sum := 0

	// Look for a * symbol
	for y, line := range lines {
		line = line + "."

		for x, character := range line {
			switch character {
			case '*':
				numbers := getNearbyNumbers(x, y, numbersMap)
				// If there are exactly 2 numbers around the gear (* symbol)
				// multiply them and add them to the sum
				if len(numbers) == 2 {
					sum += multiplyNumbers(numbers)
				}
			}
		}
	}

	fmt.Printf("Part 2: %d", sum)
}

func part1(lines []string) {
	numbersMap := generateNumbersMap(lines)

	sum := 0

	// Look for a special character
	for y, line := range lines {
		line = line + "."

		for x, character := range line {
			switch character {
			case '*', '/', '$', '+', '&', '@', '#', '%', '=', '-':
				// Get nearby numbers and add them to the sum
				numbers := getNearbyNumbers(x, y, numbersMap)
				sum += sumNumbers(numbers)
			}
		}
	}

	fmt.Printf("Part 1: %d", sum)
}

// Generate a map of coordinates where there's a number
func generateNumbersMap(lines []string) map[Coord]int {
	coordsMap := map[Coord]int{}

	for y, line := range lines {
		number := 0
		line = line + "."

		for x, character := range line {
			if digit, err := strconv.Atoi(string(character)); err == nil {
				coordsMap[Coord{x, y}] = digit
				number = number*10 + digit
			} else if number != 0 {
				// number ended, roll back to update values
				for i := x - 1; i > 0; i-- {
					if _, ok := coordsMap[Coord{i, y}]; ok {
						coordsMap[Coord{i, y}] = number
					} else {
						break
					}
				}
				number = 0
			}
		}
	}

	return coordsMap
}

// Return a slice of numbers around a coordinate
func getNearbyNumbers(x int, y int, numbersMap map[Coord]int) []int {
	numbers := []int{}

	nw, oknw := numbersMap[Coord{x - 1, y - 1}]
	n, okn := numbersMap[Coord{x, y - 1}]
	ne, okne := numbersMap[Coord{x + 1, y - 1}]
	e, oke := numbersMap[Coord{x + 1, y}]
	se, okse := numbersMap[Coord{x + 1, y + 1}]
	s, oks := numbersMap[Coord{x, y + 1}]
	sw, oksw := numbersMap[Coord{x - 1, y + 1}]
	w, okw := numbersMap[Coord{x - 1, y}]

	if oke {
		numbers = append(numbers, e)
	}
	if okw {
		numbers = append(numbers, w)
	}
	if oknw {
		numbers = append(numbers, nw)
	}
	if okn && n != nw {
		numbers = append(numbers, n)
	}
	if okne && ne != n {
		numbers = append(numbers, ne)
	}
	if oksw {
		numbers = append(numbers, sw)
	}
	if oks && s != sw {
		numbers = append(numbers, s)
	}
	if okse && se != s {
		numbers = append(numbers, se)
	}

	return numbers
}

// Sum numbers from slice
func sumNumbers(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

// Multiply numbers from slice
func multiplyNumbers(numbers []int) int {
	product := 1

	for _, number := range numbers {
		product *= number
	}

	return product
}
