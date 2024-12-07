package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/aoc-2024-go/io"
)

func main() {
	args := os.Args[1:]
	var lines []string

	if len(args) > 0 {
		lines = io.ReadFile("./sample.txt")
	} else {
		lines = io.ReadFile("./input.txt")
	}

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
	fmt.Println(part1Answer)
	fmt.Println(part2Answer)

}
