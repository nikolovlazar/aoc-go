package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/nikolovlazar/aoc-go/input"
)

func main() {
	// Define a map to convert word numbers into digits
	numberWords := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	// Read the input
	lines := input.ReadInput("/2023/day01/input.txt")

	// Define the regex pattern to match digits and number words
	pattern := regexp.MustCompile(`[0-9]{1}|one|two|three|four|five|six|seven|eight|nine`)
	numbers := []int{}
	for _, line := range lines {
		var digits []string

		// Loop until there are no matches
		for {
			// Find the index of the regex match within the single line
			indexes := pattern.FindStringIndex(line)
			if indexes == nil {
				break
			}
			// Define the start index of the match
			start := indexes[0]
			// Match the regex to obtain the value (digit or word number)
			match := pattern.FindString(line)
			// Save the value
			digits = append(digits, match)
			// Modify the line so it doesn't match the same value in the next iteration
			// We do this as a workaround for the overlapping regex matches
			// like nineighthree (nine-eight-three / 983)
			line = line[:start] + line[start+1:]
		}

		if len(digits) == 0 {
			continue
		}

		// Isolate the first and the last digit or number word
		first := digits[0]
		last := digits[len(digits)-1]

		// Check if "first" is a word number
		if v, ok := numberWords[first]; ok {
			// Convert it to a digit
			first = v
		}
		// Check if "last" is a word number
		if v, ok := numberWords[last]; ok {
			// Convert it to a digit
			last = v
		}

		// Concatenate first and last and parse to an integer
		num, err := strconv.Atoi(first + last)
		if err != nil {
			panic(err)
		}

		// Push the integer to the array
		numbers = append(numbers, num)
	}

	sum := 0

	// Sum all of the integers together
	for _, num := range numbers {
		sum = sum + num
	}

	// Profit 🌲
	fmt.Printf("The sum of all numbers is: %d", sum)
}
