package day22

import (
	"fmt"
	"strconv"

	"github.com/aoc-2024-go/io"
	"github.com/aoc-2024-go/utils"
)

var part1Answer int
var part2Answer int
var lines []string
var moduloMinusOne int = 16777216 - 1

type key = uint32
type Solution struct{}

func (s Solution) Solve(isSample bool, dirName string) (string, string) {
	var pathToInput string

	if isSample {
		pathToInput = fmt.Sprintf("./%s/sample.txt", dirName)
	} else {
		pathToInput = fmt.Sprintf("./%s/input.txt", dirName)
	}
	lines = io.ReadFile(pathToInput)

	fullPriceMap := make(map[key]int, 100000)

	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		deltas := make([]int, 2000)
		prices := make([]int, 2000)
		priceMap := make(map[key]bool, 100000)
		prevPrice := num % 10
		for j := range 2000 {
			num = nextSecret(num)
			price := num % 10
			prices[j] = price
			deltas[j] = price - prevPrice
			prevPrice = price
		}

		for j := range 2000 - 3 {
			key := toKey(j, &deltas)

			if !priceMap[key] {
				priceMap[key] = true
				fullPriceMap[key] += prices[j+3]
			}

		}
		part1Answer += num
	}
	part2Answer = 0
	for _, price := range fullPriceMap {
		part2Answer = utils.Max(part2Answer, price)
	}

	return strconv.Itoa(part1Answer), strconv.Itoa(part2Answer)
}

func nextSecret(num int) int {
	num = ((num << 6) ^ num) & (moduloMinusOne)
	num = ((num >> 5) ^ num) & (moduloMinusOne)
	num = ((num << 11) ^ num) & (moduloMinusOne)
	return num
}

func toKey(from int, deltas *[]int) key {
	key := (*deltas)[from]&31 |
		((*deltas)[from+1]&31)<<5 |
		((*deltas)[from+2]&31)<<10 |
		((*deltas)[from+3]&31)<<15
	return uint32(key)
}
