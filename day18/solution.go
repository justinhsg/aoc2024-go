package day18

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/aoc-2024-go/io"
	"github.com/aoc-2024-go/types"
)

var part1Answer int
var part2Answer string

type Pair = types.IntPair

var width, height int
var part1Threshold int
var byteTimes map[Pair]int = make(map[Pair]int)

var dirs = []Pair{
	types.NewIntPair(-1, 0),
	types.NewIntPair(0, 1),
	types.NewIntPair(1, 0),
	types.NewIntPair(0, -1),
}

var ufds map[Pair]Pair = make(map[Pair]Pair)

func parent(node Pair) Pair {
	if ufds[node] == node {
		return node
	}
	ufds[node] = parent(ufds[node])
	return ufds[node]
}

func join(node1 Pair, node2 Pair) {
	ufds[parent(node1)] = parent(node2)
}

func isSameSet(node1 Pair, node2 Pair) bool {
	return parent(node1) == parent(node2)
}

type Solution struct{}

func (s Solution) Solve(isSample bool, dirName string) (string, string) {
	var pathToInput string

	if isSample {
		pathToInput = fmt.Sprintf("./%s/sample.txt", dirName)
		width, height = 7, 7
		part1Threshold = 12
	} else {
		pathToInput = fmt.Sprintf("./%s/input.txt", dirName)
		width, height = 71, 71
		part1Threshold = 1024
	}
	lines := io.ReadFile(pathToInput)

	var bytes []Pair = make([]Pair, len(lines))
	for time, line := range lines {
		parts := strings.Split(line, ",")
		byteCol, _ := strconv.Atoi(parts[0])
		byteRow, _ := strconv.Atoi(parts[1])
		bytePair := Pair{Fst: byteRow, Snd: byteCol}
		byteTimes[bytePair] = time
		bytes[time] = bytePair
	}

	part1Answer, _ = pathFind(part1Threshold - 1)
	for i := range height {
		for j := range width {
			ufds[types.NewIntPair(i, j)] = types.NewIntPair(i, j)
		}
	}
	floodFill()
	start := types.NewIntPair(0, 0)
	end := types.NewIntPair(height-1, width-1)
	for i := len(bytes) - 1; i >= 0; i-- {
		bytePos := bytes[i]
		for _, dir := range dirs {
			adjPos := types.AddPair(bytePos, dir)
			if adjPos.Fst >= 0 && adjPos.Fst < height && adjPos.Snd >= 0 && adjPos.Snd < width {
				time, isByte := byteTimes[adjPos]
				if !isByte || time >= i {
					join(bytePos, adjPos)
				}
			}
		}
		if isSameSet(start, end) {
			part2Answer = fmt.Sprintf("%d,%d", bytePos.Snd, bytePos.Fst)
			break
		}
	}

	return strconv.Itoa(part1Answer), part2Answer
}

func floodFill() {
	var visited [][]bool = make([][]bool, height)
	for i := range height {
		visited[i] = make([]bool, width)
	}

	var toVisit []Pair

	for i := range height {
		for j := range width {
			if visited[i][j] {
				continue
			}
			root := types.NewIntPair(i, j)
			toVisit = append(toVisit, root)
			visited[i][j] = true
			for len(toVisit) != 0 {
				curPos := toVisit[0]
				toVisit = toVisit[1:]
				for _, dir := range dirs {
					newPos := types.AddPair(curPos, dir)
					if newPos.Fst >= 0 && newPos.Fst < height && newPos.Snd >= 0 && newPos.Snd < width {
						_, hasByte := byteTimes[newPos]
						if !hasByte && !visited[newPos.Fst][newPos.Snd] {
							visited[newPos.Fst][newPos.Snd] = true
							join(newPos, root)
							toVisit = append(toVisit, newPos)
						}
					}
				}
			}
		}
	}
}

func pathFind(time int) (int, bool) {
	var distances [][]int = make([][]int, height)
	for i := range height {

		distances[i] = make([]int, width)

		for j := range width {
			distances[i][j] = math.MaxInt64
		}
	}

	var toVisit []Pair
	distances[0][0] = 0

	toVisit = append(toVisit, types.NewIntPair(0, 0))

	for len(toVisit) != 0 {

		curPos := toVisit[0]
		curDist := distances[curPos.Fst][curPos.Snd]
		toVisit = toVisit[1:]
		for _, dir := range dirs {
			newPos := types.AddPair(curPos, dir)
			if newPos.Fst >= 0 && newPos.Fst < height && newPos.Snd >= 0 && newPos.Snd < width {
				byteTime, hasByte := byteTimes[newPos]
				if (!hasByte || byteTime > time) && distances[newPos.Fst][newPos.Snd] > curDist+1 {
					distances[newPos.Fst][newPos.Snd] = curDist + 1
					toVisit = append(toVisit, newPos)
				}
			}
		}
	}
	if distances[height-1][width-1] == math.MaxInt64 {
		return -1, false
	} else {
		return distances[height-1][width-1], true
	}
}
