package day19

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aoc-2024-go/io"
	"github.com/aoc-2024-go/types"
)

var part1Answer, part2Answer int

type Pair = types.IntPair

var fragments map[byte][]string = make(map[byte][]string)

type Solution struct{}

func (s Solution) Solve(isSample bool, dirName string) (string, string) {
	var pathToInput string

	if isSample {
		pathToInput = fmt.Sprintf("./%s/sample.txt", dirName)
	} else {
		pathToInput = fmt.Sprintf("./%s/input.txt", dirName)
	}
	lines := io.ReadFile(pathToInput)

	fragmentArr := strings.Split(lines[0], ", ")
	for _, fragment := range fragmentArr {
		firstRune := fragment[0]
		fragments[firstRune] = append(fragments[firstRune], fragment)
	}
	designs := lines[2:]
	for _, design := range designs {
		ways := tryDesign(design)
		if ways > 0 {
			part1Answer += 1
			part2Answer += ways
		}

	}
	return strconv.Itoa(part1Answer), strconv.Itoa(part2Answer)
}

func tryDesign(design string) int {
	var canDesign []int = make([]int, len(design)+1)

	canDesign[0] = 1

	for i := range len(design) {
		if canDesign[i] > 0 {
			for _, candidate := range fragments[design[i]] {
				candidateLength := len(candidate)

				// fmt.Println(candidate, design[i:i+candidateLength])
				if i+candidateLength <= len(design) && design[i:i+candidateLength] == candidate {
					canDesign[i+candidateLength] += canDesign[i]
				}
			}
		}
	}
	return canDesign[len(design)]

}
