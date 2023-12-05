package utils

import (
	"strconv"
	"strings"
)

func ExtractNumbers(value string) []int {
	value = strings.TrimSpace(value)
	numbers := []int{}
	for _, number := range strings.Split(value, " ") {
		parsed, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, parsed)
	}

	return numbers
}
