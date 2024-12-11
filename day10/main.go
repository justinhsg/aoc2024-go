package main

import (
	"fmt"
	"os"

	"github.com/aoc-2024-go/io"
	"github.com/aoc-2024-go/types"
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

	var heights [][]types.IntPair = make([][]types.IntPair, 10)
	var nPaths [][]int = make([][]int, len(lines))
	var grid [][]int = make([][]int, len(lines))

	for row, line := range lines {
		grid[row] = make([]int, len(line))
		nPaths[row] = make([]int, len(line))
		for col := range line {
			curPos := types.NewIntPair(row, col)
			height := int(line[col] - '0')
			grid[row][col] = height
			heights[height] = append(heights[height], curPos)

		}
	}
	var h, w int = len(grid), len(grid[0])
	var dirs []types.IntPair = []types.IntPair{
		types.NewIntPair(-1, 0),
		types.NewIntPair(0, 1),
		types.NewIntPair(1, 0),
		types.NewIntPair(0, -1)}
	for _, start := range heights[0] {
		var trails map[types.IntPair]bool = make(map[types.IntPair]bool)
		var visited map[types.IntPair]bool = make(map[types.IntPair]bool)
		var toVisit []types.IntPair = []types.IntPair{start}
		for len(toVisit) > 0 {
			curPos := toVisit[0]
			toVisit = toVisit[1:]
			curVal := grid[curPos.Fst][curPos.Snd]
			if curVal == 9 {
				trails[curPos] = true
				continue
			}
			visited[curPos] = true
			for _, dir := range dirs {
				newPos := types.AddPair(curPos, dir)
				if newPos.Fst >= 0 && newPos.Fst < h && newPos.Snd >= 0 && newPos.Snd < w {
					if !visited[newPos] && grid[newPos.Fst][newPos.Snd] == curVal+1 {
						toVisit = append(toVisit, newPos)
					}
				}
			}
		}
		part1Answer += len(trails)
	}

	for _, ends := range heights[9] {
		nPaths[ends.Fst][ends.Snd] = 1
	}

	for height := 8; height >= 0; height-- {
		for _, curPos := range heights[height] {
			for _, dir := range dirs {
				newPos := types.AddPair(curPos, dir)

				if newPos.Fst >= 0 && newPos.Fst < h && newPos.Snd >= 0 && newPos.Snd < w {
					if grid[newPos.Fst][newPos.Snd] == height+1 {
						nPaths[curPos.Fst][curPos.Snd] += nPaths[newPos.Fst][newPos.Snd]
					}
				}
			}
		}
	}

	for _, starts := range heights[0] {
		part2Answer += nPaths[starts.Fst][starts.Snd]
	}
	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
}
