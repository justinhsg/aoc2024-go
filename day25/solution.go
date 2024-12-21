package day25

import (
	"fmt"
	"strconv"

	"github.com/aoc-2024-go/io"
)

var part1Answer, part2Answer int
var lines []string

type Solution struct{}

func (s Solution) Solve(isSample bool, dirName string) (string, string) {
	var pathToInput string

	if isSample {
		pathToInput = fmt.Sprintf("./%s/sample.txt", dirName)
	} else {
		pathToInput = fmt.Sprintf("./%s/input.txt", dirName)
	}
	lines = io.ReadFile(pathToInput)

	return strconv.Itoa(part1Answer), "Merry Christmas!"
}
