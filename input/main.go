package input

import (
	"bufio"
	"os"

	"github.com/nikolovlazar/aoc-go/utils"
)

func ReadInput(filename string) []string {
	filePath := utils.GetFileDir(filename)

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}
