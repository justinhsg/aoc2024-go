package main

import (
	"fmt"
	"maps"
	"os"

	"github.com/aoc-2024-go/io"
	"github.com/aoc-2024-go/types"
)

var start types.IntPair
var width, height int
var obstacles map[types.IntPair]bool = make(map[types.IntPair]bool)
var dRow, dCol []int = []int{-1, 0, 1, 0}, []int{0, 1, 0, -1}

func main() {
	args := os.Args[1:]
	var lines []string

	if len(args) > 0 {
		lines = io.ReadFile("./sample.txt")
	} else {
		lines = io.ReadFile("./input.txt")
	}

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
			}
		}
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

	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
}

func tryWithObstacle(extra types.IntPair) bool {
	visited := make(map[types.IntTriple]bool)

	posAndDir := types.NewIntTriple(start.Fst, start.Snd, 0)
	isLoop := false

	for (posAndDir.Fst >= 0 && posAndDir.Fst < height && posAndDir.Snd >= 0 && posAndDir.Snd < width) &&
		!isLoop {
		if visited[posAndDir] {
			isLoop = true
			break
		}
		visited[posAndDir] = true
		nextPos := types.NewIntPair(posAndDir.Fst+dRow[posAndDir.Thd], posAndDir.Snd+dCol[posAndDir.Thd])
		if obstacles[nextPos] || nextPos == extra {
			posAndDir = types.NewIntTriple(posAndDir.Fst, posAndDir.Snd, (posAndDir.Thd+1)%4)
		} else {
			posAndDir = types.NewIntTriple(nextPos.Fst, nextPos.Snd, posAndDir.Thd)
		}
	}
	return isLoop

}
