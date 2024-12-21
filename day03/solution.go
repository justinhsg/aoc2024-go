package day03

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/aoc-2024-go/io"
)

type Solution struct{}

func (s Solution) Solve(isSample bool, dirName string) (string, string) {
	var pathToInput string

	if isSample {
		pathToInput = fmt.Sprintf("./%s/sample.txt", dirName)
	} else {
		pathToInput = fmt.Sprintf("./%s/input.txt", dirName)
	}

	lines := io.ReadFile(pathToInput)

	mulPattern, _ := regexp.Compile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)

	part1Answer := 0
	part2Answer := 0
	isEnabled := true
	for _, line := range lines {
		matches := mulPattern.FindAllString(line, -1)

		for _, match := range matches {
			var (
				int1 int
				int2 int
			)
			if match == "do()" {
				isEnabled = true
			} else if match == "don't()" {
				isEnabled = false
			} else {
				fmt.Sscanf(match, "mul(%d,%d)", &int1, &int2)
				part1Answer += int1 * int2
				if isEnabled {
					part2Answer += int1 * int2
				}
			}

		}
	}
	return strconv.Itoa(part1Answer), strconv.Itoa(part2Answer)

}
