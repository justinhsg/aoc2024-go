package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/aoc-2024-go/io"
	"github.com/aoc-2024-go/types"
	"github.com/aoc-2024-go/utils"
)

var part1Answer, part2Answer string

type Triple = types.IntTriple
type Pair = types.IntPair
type PriorityQueue = types.PriorityQueue[Triple]
type Item = types.PriorityQueueItem[Triple]

func main() {
	args := os.Args[1:]
	var lines []string

	if len(args) == 1 {
		lines = io.ReadFile("./sample.txt")

	} else {
		lines = io.ReadFile("./input.txt")
	}

	var origA, origB, origC int
	fmt.Sscanf(lines[0], "Register A: %d", &origA)
	fmt.Sscanf(lines[1], "Register B: %d", &origB)
	fmt.Sscanf(lines[2], "Register C: %d", &origC)

	var program []int = utils.Map(strings.Split(strings.Split(lines[4], " ")[1], ","), func(x string) int {
		i, _ := strconv.Atoi(x)
		return i
	})

	part1Output := runProg(program, origA, true)
	part1Answer = strings.Join(utils.Map(part1Output, func(i int) string {
		return strconv.Itoa(i)
	}), ",")

	var candidates []int = []int{0}
	part2Answer := 0
	for i := len(program) - 1; i >= 0; i-- {
		target := program[i]
		var newCandidates []int
		for _, candidate := range candidates {
			for test := 0; test < 8; test++ {
				tOUt := runProg(program, candidate*8+test, false)
				if tOUt[0] == target {
					newCandidates = append(newCandidates, candidate*8+test)
				}
			}
		}
		candidates = newCandidates
	}
	part2Answer = math.MaxInt64
	for _, candidate := range candidates {
		if candidate < part2Answer {
			part2Answer = candidate
		}
	}
	fmt.Println(part1Answer)
	fmt.Println(part2Answer)

}

func evalCombo(operand, a, b, c int) int {
	if operand <= 3 {
		return operand
	}
	if operand == 4 {
		return a
	}
	if operand == 5 {
		return b
	}
	if operand == 6 {
		return c
	}
	return 0
}

func runProg(program []int, initA int, loop bool) []int {
	var a, b, c = initA, 0, 0
	var output []int
	for i := 0; i != len(program); i += 2 {
		opCode := program[i]
		literal := program[i+1]
		combo := evalCombo(literal, a, b, c)
		switch opCode {
		case 0:

			result := a / (1 << combo)
			a = result
		case 1:
			result := b ^ literal
			b = result
		case 2:
			b = combo % 8
		case 3:
			if a != 0 && loop {
				i = literal - 2
			}
		case 4:
			result := b ^ c
			b = result
		case 5:
			output = append(output, combo%8)
		case 6:
			result := a / (1 << combo)
			b = result
		case 7:
			result := a / (1 << combo)
			c = result
		}
	}
	return output
}
