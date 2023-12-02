package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nikolovlazar/aoc-go/input"
)

func main() {
	// Part 1: max cubes per color
	// config := map[string]int{
	// 	"red":   12,
	// 	"green": 13,
	// 	"blue":  14,
	// }

	// Read the input
	lines := input.ReadInput("/2023/day02/input.txt")

	sum := 0

	for _, game_line := range lines {
		game_records := strings.Split(game_line, ": ")
		records := game_records[1]
		record_parts := strings.Split(records, " ")
		// Part 1: impossible game flag
		// impossible := false
		min_map := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for i := 0; i < len(record_parts); i += 2 {
			number_str, color := record_parts[i], record_parts[i+1]
			number, err := strconv.Atoi(number_str)
			if err != nil {
				panic(err)
			}

			var color_name string
			if last_letter := color[len(color)-1]; last_letter == ',' || last_letter == ';' {
				color_name = color[:len(color)-1]
			} else {
				color_name = color
			}

			min_map[color_name] = max(min_map[color_name], number)

			// Part 1: mark game impossible
			// if number > config[color_name] {
			// 	// This game is impossible
			// 	impossible = true
			// 	break
			// }
		}

		power := min_map["red"] * min_map["green"] * min_map["blue"]
		sum += power

		// Part 1: sum if game is possible
		// if !impossible {
		// 	sum += index + 1
		// }
	}

	fmt.Println(sum)
}
