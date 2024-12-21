package day08

import (
	"fmt"
	"strconv"

	"github.com/aoc-2024-go/io"
	pair "github.com/aoc-2024-go/types"
)

var part1Answer, part2Answer int

type Solution struct{}

func (s Solution) Solve(isSample bool, dirName string) (string, string) {
	var pathToInput string

	if isSample {
		pathToInput = fmt.Sprintf("./%s/sample.txt", dirName)
	} else {
		pathToInput = fmt.Sprintf("./%s/input.txt", dirName)
	}
	lines := io.ReadFile(pathToInput)

	width := len(lines[0])
	height := len(lines)

	satellites := make(map[rune][]pair.IntPair)

	for row, line := range lines {
		for col, r := range line {
			if r != rune('.') {
				// if satellites[r] == nil {
				// 	satellites[r] = make([]Pair)
				// }
				satellites[r] = append(satellites[r], pair.NewIntPair(row, col))
			}
		}
	}

	antiNodes := make(map[pair.IntPair]bool)
	harmNodes := make(map[pair.IntPair]bool)
	for _, locations := range satellites {
		for i, p1 := range locations {
			for j, p2 := range locations {
				if i == j {
					continue
				}
				dir := pair.DiffPair(p2, p1)
				anti := pair.AddPair(p2, dir)
				if anti.Fst >= 0 && anti.Fst < height && anti.Snd >= 0 && anti.Snd < width {
					antiNodes[anti] = true
				}
				harm := p2
				for harm.Fst >= 0 && harm.Fst < height && harm.Snd >= 0 && harm.Snd < width {
					harmNodes[harm] = true
					harm = pair.AddPair(harm, dir)
				}
			}
		}
	}
	part1Answer = len(antiNodes)
	part2Answer = len(harmNodes)
	return strconv.Itoa(part1Answer), strconv.Itoa(part2Answer)
}
