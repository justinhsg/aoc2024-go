package day21

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/aoc-2024-go/io"
	"github.com/aoc-2024-go/types"
)

var part1Answer, part2Answer int

type Pair = types.IntPair
type Path = types.Pair[byte, byte]

var dirLocs = map[byte]Pair{
	'^': {Fst: 0, Snd: 1},
	'A': {Fst: 0, Snd: 2},
	'<': {Fst: 1, Snd: 0},
	'v': {Fst: 1, Snd: 1},
	'>': {Fst: 1, Snd: 2},
}

var numLocs = map[byte]Pair{
	'7': {Fst: 0, Snd: 0},
	'8': {Fst: 0, Snd: 1},
	'9': {Fst: 0, Snd: 2},
	'4': {Fst: 1, Snd: 0},
	'5': {Fst: 1, Snd: 1},
	'6': {Fst: 1, Snd: 2},
	'1': {Fst: 2, Snd: 0},
	'2': {Fst: 2, Snd: 1},
	'3': {Fst: 2, Snd: 2},
	'0': {Fst: 3, Snd: 1},
	'A': {Fst: 3, Snd: 2},
}

var numAdjList = map[byte]string{
	'7': "X84X",
	'8': "X957",
	'9': "XX68",
	'4': "751X",
	'5': "8624",
	'6': "9X35",
	'1': "42XX",
	'2': "5301",
	'3': "6XA2",
	'0': "2AXX",
	'A': "3XX0",
}

var dirAdjList = map[byte]string{
	'^': "XAvX",
	'A': "XX>^",
	'<': "XvXX",
	'v': "^>X<",
	'>': "AXXv",
}
var numbers = []byte{'A', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
var dirs = []byte{'^', '>', 'v', '<', 'A'}
var dirPaths map[byte]map[byte][]string = map[byte]map[byte][]string{}
var numPaths map[byte]map[byte][]string = map[byte]map[byte][]string{}

type Solution struct{}

func (s Solution) Solve(isSample bool, dirName string) (string, string) {
	var pathToInput string

	if isSample {
		pathToInput = fmt.Sprintf("./%s/sample.txt", dirName)
	} else {
		pathToInput = fmt.Sprintf("./%s/input.txt", dirName)
	}
	lines := io.ReadFile(pathToInput)
	for _, from := range numbers {
		numPaths[from] = make(map[byte][]string)
		for _, to := range numbers {
			numPaths[from][to] = getPaths(from, to, &numLocs, &numAdjList)
		}
	}

	for _, from := range dirs {
		dirPaths[from] = make(map[byte][]string)
		for _, to := range dirs {
			dirPaths[from][to] = getPaths(from, to, &dirLocs, &dirAdjList)
		}
	}

	for _, line := range lines {
		augLine := "A" + line
		totMovesPart1 := 0
		totMovesPart2 := 0
		for i := range len(augLine) - 1 {
			from, to := augLine[i], augLine[i+1]
			bestMovePart1 := math.MaxInt64
			bestMovePart2 := math.MaxInt64
			for _, numPath := range numPaths[from][to] {
				curMovePart1 := calcMoves(numPath, 1)

				if curMovePart1 < bestMovePart1 {
					bestMovePart1 = curMovePart1
				}
				curMovePart2 := calcMoves(numPath, 24)
				if curMovePart2 < bestMovePart2 {
					bestMovePart2 = curMovePart2
				}
			}

			totMovesPart1 += bestMovePart1
			totMovesPart2 += bestMovePart2
		}
		numPart, _ := strconv.Atoi(line[:len(line)-1])
		part1Answer += numPart * totMovesPart1
		part2Answer += numPart * totMovesPart2
	}
	return strconv.Itoa(part1Answer), strconv.Itoa(part2Answer)
}

func getPaths(from byte, to byte, locs *map[byte]Pair, adjList *map[byte]string) []string {
	fromPos, toLocs := (*locs)[from], (*locs)[to]
	down, right := types.DiffPair(toLocs, fromPos).Destruct()
	var hString, vString string
	if down < 0 {
		vString = strings.Repeat("^", -down)
	} else {
		vString = strings.Repeat("v", down)
	}
	if right < 0 {
		hString = strings.Repeat("<", -right)
	} else {
		hString = strings.Repeat(">", right)
	}
	if down == 0 || right == 0 {
		return []string{vString + hString + "A"}
	} else {
		options := []string{}
		option1 := vString + hString + "A"
		if verify(from, to, option1, adjList) {
			options = append(options, option1)
		}
		option2 := hString + vString + "A"
		if verify(from, to, option2, adjList) {
			options = append(options, option2)
		}
		return options
	}
}

var dirToIndex = map[byte]int{'^': 0, '>': 1, 'v': 2, '<': 3}

func verify(from byte, to byte, path string, adjList *map[byte]string) bool {
	curByte := from
	for _, dir := range path[:len(path)-1] {
		curByte = (*adjList)[curByte][dirToIndex[byte(dir)]]
		if curByte == 'X' {
			return false
		}
	}
	return curByte == to && path[len(path)-1] == 'A'
}

var memo map[string]map[int]int = make(map[string]map[int]int)

func calcMoves(target string, nLayers int) int {
	nestedMap, inMemo := memo[target]
	if inMemo {
		prevAnswer, seenBefore := nestedMap[nLayers]
		if seenBefore {
			return prevAnswer
		}
	} else {
		memo[target] = make(map[int]int)
	}

	augTarget := "A" + target
	nMoves := 0
	for i := range len(augTarget) - 1 {
		from, to := augTarget[i], augTarget[i+1]
		options := dirPaths[from][to]
		if nLayers == 0 {
			nMoves += len(options[0])
		} else {
			bestMoves := math.MaxInt64
			for _, option := range options {
				curMove := calcMoves(option, nLayers-1)
				if curMove < bestMoves {
					bestMoves = curMove
				}
			}
			nMoves += bestMoves
		}
	}
	memo[target][nLayers] = nMoves
	return nMoves
}
