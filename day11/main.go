package main

import (
	"fmt"
	"maps"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/aoc-2024-go/io"
	"github.com/aoc-2024-go/utils"
)

var part1Answer, part2Answer int

func main() {
	args := os.Args[1:]
	var lines []string

	if len(args) > 0 {
		lines = io.ReadFile("./sample.txt")
	} else {
		lines = io.ReadFile("./input.txt")
	}

	var line string = lines[0]
	var stones []int = utils.Map(strings.Split(line, " "), func(str string) int {
		i, _ := strconv.Atoi(str)
		return i
	})

	var memo map[int][]int = make(map[int][]int)

	var curStones map[int]int = make(map[int]int)
	for _, stone := range stones {
		curStones[stone] = 1
	}
	for blink := 1; blink <= 75; blink++ {
		var nextStones map[int]int = make(map[int]int)
		for stone := range maps.Keys(curStones) {
			stoneQty := curStones[stone]
			if len(memo[stone]) == 0 {
				memo[stone] = calcNext(stone)
			}
			for _, nextStone := range memo[stone] {
				nextStones[nextStone] += stoneQty
			}
		}

		curStones = nextStones
		if blink == 25 {
			part1Answer = countStones(&curStones)
		}
		if blink == 75 {
			part2Answer = countStones(&curStones)
		}
	}
	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
}

func countStones(stoneMap *map[int]int) int {
	count := 0
	for stone := range maps.Keys(*stoneMap) {
		count += (*stoneMap)[stone]
	}
	return count
}

func calcNext(stone int) []int {
	if stone == 0 {
		return []int{1}
	} else if nDigits := utils.NDigits(stone); nDigits%2 == 0 {
		pow10 := int(math.Pow10(nDigits / 2))
		return []int{stone / pow10, stone % pow10}
	} else {
		return []int{stone * 2024}
	}
}
