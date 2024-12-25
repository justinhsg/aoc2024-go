package day25

import (
	"fmt"
	"strconv"

	"github.com/aoc-2024-go/io"
)

var part1Answer int
var sections [][]string

var allCodes []int64

type Solution struct{}

func (s Solution) Solve(isSample bool, dirName string) (string, string) {
	var pathToInput string

	if isSample {
		pathToInput = fmt.Sprintf("./%s/sample.txt", dirName)
	} else {
		pathToInput = fmt.Sprintf("./%s/input.txt", dirName)
	}
	sections = io.SplitIntoSections(pathToInput)
	allCodes = make([]int64, len(sections))
	for i, section := range sections {

		allCodes[i] = encode(&section)
	}

	for i, codeA := range allCodes {
		for _, codeB := range allCodes[i+1:] {
			if codeA&codeB == 0 {
				part1Answer += 1
			}
		}
	}
	return strconv.Itoa(part1Answer), "Merry Christmas!"
}

func encode(grid *[]string) int64 {
	var code int64 = 0
	for _, line := range *grid {
		for _, c := range line {
			if c == rune('#') {
				code |= 1
			}
			code <<= 1
		}
	}
	return code
}
