package day06

import (
	"fmt"
	"maps"
	"slices"
	"strconv"

	"github.com/aoc-2024-go/io"
	"github.com/aoc-2024-go/types"
	"github.com/aoc-2024-go/utils"
)

var start types.IntPair
var width, height int
var obstacles map[types.IntPair]bool = make(map[types.IntPair]bool)
var dRow, dCol []int = []int{-1, 0, 1, 0}, []int{0, 1, 0, -1}
var obstaclesByRow, obstaclesByColumn map[int][]int = make(map[int][]int), make(map[int][]int)

type Solution struct{}

func (s Solution) Solve(isSample bool, dirName string) (string, string) {
	var pathToInput string

	if isSample {
		pathToInput = fmt.Sprintf("./%s/sample.txt", dirName)
	} else {
		pathToInput = fmt.Sprintf("./%s/input.txt", dirName)
	}
	lines := io.ReadFile(pathToInput)

	width = len(lines[0])
	height = len(lines)
	var pos types.IntPair
	visited := make(map[types.IntPair]bool)
	dir := 0

	for row, line := range lines {
		for col, r := range line {
			if r == rune('^') {
				pos = types.NewIntPair(row, col)
				start = types.NewIntPair(row, col)
			}
			if r == rune('#') {
				obstacles[types.NewIntPair(row, col)] = true
				obstaclesByColumn[col] = append(obstaclesByColumn[col], row)
				obstaclesByRow[row] = append(obstaclesByRow[row], col)
			}
		}
	}

	for col := range maps.Keys(obstaclesByColumn) {
		slices.Sort(obstaclesByColumn[col])
	}
	for row := range maps.Keys(obstaclesByRow) {
		slices.Sort(obstaclesByRow[row])
	}

	for pos.Fst >= 0 && pos.Fst < height && pos.Snd >= 0 && pos.Snd < width {
		visited[pos] = true
		nextPos := types.NewIntPair(pos.Fst+dRow[dir], pos.Snd+dCol[dir])
		if obstacles[nextPos] {
			dir = (dir + 1) % 4
		} else {
			pos = nextPos
		}
	}

	var part1Answer, part2Answer int
	for extra := range maps.Keys(visited) {
		if extra == start {
			continue
		}
		if tryWithObstacle(extra) {
			part2Answer += 1
		}
	}

	part1Answer = len(visited)

	return strconv.Itoa(part1Answer), strconv.Itoa(part2Answer)
}

func tryWithObstacle(extra types.IntPair) bool {
	originalRowObs := obstaclesByRow[extra.Fst]
	obstaclesByRow[extra.Fst] = insertIntoArray(originalRowObs, extra.Snd)

	originalColObs := obstaclesByColumn[extra.Snd]
	obstaclesByColumn[extra.Snd] = insertIntoArray(originalColObs, extra.Fst)

	visited := make(map[types.IntTriple]bool)

	posAndDir := types.NewIntTriple(start.Fst, start.Snd, 0)
	isLoop := false

	for !isLoop {
		if visited[posAndDir] {
			isLoop = true
			break
		}
		visited[posAndDir] = true
		row := posAndDir.Fst
		col := posAndDir.Snd
		dir := posAndDir.Thd

		isObstructed := false
		if dir == 0 {
			for idx := len(obstaclesByColumn[col]) - 1; idx >= 0; idx-- {
				if obstaclesByColumn[col][idx] < row {
					row = obstaclesByColumn[col][idx] + 1
					isObstructed = true
					break
				}
			}
		} else if dir == 1 {
			for idx := 0; idx < len(obstaclesByRow[row]); idx++ {
				if obstaclesByRow[row][idx] > col {
					col = obstaclesByRow[row][idx] - 1
					isObstructed = true
					break
				}
			}
		} else if dir == 2 {
			for idx := 0; idx < len(obstaclesByColumn[col]); idx++ {
				if obstaclesByColumn[col][idx] > row {
					row = obstaclesByColumn[col][idx] - 1
					isObstructed = true
					break
				}
			}
		} else if dir == 3 {
			for idx := len(obstaclesByRow[row]) - 1; idx >= 0; idx-- {
				if obstaclesByRow[row][idx] < col {
					col = obstaclesByRow[row][idx] + 1
					isObstructed = true
					break
				}
			}
		}
		if !isObstructed {
			break
		}
		posAndDir = types.NewIntTriple(row, col, (dir+1)%4)
	}
	obstaclesByRow[extra.Fst] = originalRowObs
	obstaclesByColumn[extra.Snd] = originalColObs
	return isLoop
}

func insertIntoArray(arr []int, i int) []int {
	newArr := make([]int, len(arr)+1)
	idx, _ := utils.Find(arr, func(x int) bool {
		return x > i
	})
	if idx == -1 {
		idx = len(arr)
	}
	copy(newArr[:idx], arr[:idx])
	newArr[idx] = i
	copy(newArr[idx+1:], arr[idx:])
	return newArr
}
