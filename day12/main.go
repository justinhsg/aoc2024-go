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

	var width, height int = len(lines[0]), len(lines)
	var visited [][]bool = make([][]bool, height)

	for i := 0; i < height; i++ {
		visited[i] = make([]bool, width)
	}

	for row, line := range lines {
		for col := range line {
			if !visited[row][col] {
				part1Cost, part2Cost := visit(row, col, &visited, &lines, width, height)
				part1Answer += part1Cost
				part2Answer += part2Cost
			}
		}
	}

	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
}

var dirs []types.IntPair = []types.IntPair{
	types.NewIntPair(-1, 0),
	types.NewIntPair(0, 1),
	types.NewIntPair(1, 0),
	types.NewIntPair(0, -1)}

func visit(row int, col int, visited *[][]bool, lines *[]string, width int, height int) (int, int) {
	var toVisit []types.IntPair
	start := types.NewIntPair(row, col)
	toVisit = append(toVisit, start)
	(*visited)[row][col] = true
	plant := (*lines)[row][col]

	var plants map[types.IntPair]bool = make(map[types.IntPair]bool)
	plants[start] = true
	var area, perimeter int = 0, 0
	var edges [][]types.IntPair = make([][]types.IntPair, 4)
	for len(toVisit) > 0 {
		curPos := toVisit[0]
		toVisit = toVisit[1:]
		area += 1
		(*visited)[curPos.Fst][curPos.Snd] = true
		tSides := 4
		for idx, dir := range dirs {
			newPos := types.AddPair(dir, curPos)
			if newPos.Fst >= 0 && newPos.Fst < height && newPos.Snd >= 0 && newPos.Snd < width {
				newPlant := (*lines)[newPos.Fst][newPos.Snd]

				if newPlant == plant {
					tSides -= 1
					if !(*visited)[newPos.Fst][newPos.Snd] {
						(*visited)[newPos.Fst][newPos.Snd] = true
						toVisit = append(toVisit, newPos)
						plants[newPos] = true
					}
				} else {
					edges[idx] = append(edges[idx], curPos)
				}
			} else {
				edges[idx] = append(edges[idx], curPos)
			}
		}
		perimeter += tSides
	}
	sides := nCorners(&plants)
	return area * perimeter, area * sides
}

func nCorners(allPos *map[types.IntPair]bool) int {
	nSides := 0
	for pos := range *allPos {
		for i, dir1 := range dirs {
			dir2 := dirs[(i+1)%4]
			adjPos1 := types.AddPair(pos, dir1)
			adjPos2 := types.AddPair(pos, dir2)
			diagPos := types.AddPair(adjPos1, dir2)

			diagPosIsOutside := !(*allPos)[diagPos]
			adjPosAreInside := ((*allPos)[adjPos1] && (*allPos)[adjPos2])
			adjPosAreOutside := !(*allPos)[adjPos1] && !(*allPos)[adjPos2]
			if diagPosIsOutside && (adjPosAreInside || adjPosAreOutside) {
				nSides += 1
			}
			if !diagPosIsOutside && adjPosAreOutside {
				nSides += 1
			}
		}
	}
	return nSides
}
