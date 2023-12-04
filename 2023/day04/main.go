package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nikolovlazar/aoc-go/input"
)

func main() {
	// Read the input
	lines := input.ReadInput("/2023/day04/input.txt")

	part1(lines)
	part2(lines)
}

func part2(lines []string) {
	cards := map[int]int{}

	for index, line := range lines {
		parts := strings.Split(line, ": ")
		numbers := strings.Split(parts[1], " | ")
		winningNumbers, myNumbers := extractNumbers(numbers[0]), extractNumbers(numbers[1])
		hits := 0

		for _, winningNumber := range winningNumbers {
			for _, myNumber := range myNumbers {
				if winningNumber == myNumber {
					hits += 1

					break
				}
			}
		}

		cards[index+1] = hits
	}

	copies := map[int]int{}

	for card := 1; card <= len(cards); card++ {
		value := cards[card]
		numCopies := copies[card]

		for i := card + 1; i <= card+value; i++ {
			if _, ok := cards[i]; ok {
				copies[i] += numCopies + 1
			} else {
				break
			}
		}
	}

	sum := 0

	for card := 1; card <= len(cards); card++ {
		value := copies[card] + 1

		sum += value
	}

	fmt.Printf("Part 2: %d", sum)
}

func part1(lines []string) {
	points := 0

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		numbers := strings.Split(parts[1], " | ")
		winningNumbers, myNumbers := extractNumbers(numbers[0]), extractNumbers(numbers[1])
		hits := 0

		for _, winningNumber := range winningNumbers {
			for _, myNumber := range myNumbers {
				if winningNumber == myNumber {
					if hits == 0 {
						hits += 1
					} else {
						hits += hits
					}

					break
				}
			}
		}

		points += hits
	}

	fmt.Printf("Part 1: %d", points)
}

func extractNumbers(card string) []int {
	card = strings.ReplaceAll(card, "  ", " ")
	card = strings.TrimSpace(card)
	numbers := []int{}
	for _, number := range strings.Split(card, " ") {
		parsed, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, parsed)
	}

	return numbers
}
