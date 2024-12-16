package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"

	"github.com/aoc-2024-go/io"
	"github.com/aoc-2024-go/types"
)

var part1Answer, part2Answer int

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

	width := len(lines[0])
	height := len(lines)

	grid := make([][]byte, height)
	distances := make([][][]int, height)
	fromStates := make([][][][]Triple, height)
	var start Triple
	var end Pair
	var dirs = []Pair{
		types.NewIntPair(-1, 0),
		types.NewIntPair(0, 1),
		types.NewIntPair(1, 0),
		types.NewIntPair(0, -1),
	}
	for row, line := range lines {
		grid[row] = make([]byte, width)
		distances[row] = make([][]int, width)
		fromStates[row] = make([][][]Triple, width)
		for col, ch := range line {
			grid[row][col] = byte(ch)
			distances[row][col] = []int{math.MaxInt32, math.MaxInt32, math.MaxInt32, math.MaxInt32}
			fromStates[row][col] = make([][]Triple, 4)
			if ch == 'S' {
				start = types.NewIntTriple(row, col, 1)
				grid[row][col] = '.'
				distances[row][col][1] = 0
			}
			if ch == 'E' {
				end = types.NewIntPair(row, col)
				grid[row][col] = '.'
			}

		}
	}

	var pq PriorityQueue = make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{
		Value:    start,
		Priority: 0,
	})

	var deltaDirs []int = []int{0, 3, 1}
	for pq.Len() > 0 {
		curItem := heap.Pop(&pq).(*Item)
		row, col, dir := curItem.Value.Fst, curItem.Value.Snd, curItem.Value.Thd
		curDist := curItem.Priority

		if curDist > distances[row][col][dir] {
			continue
		}
		for _, deltaDir := range deltaDirs {
			newDir := (dir + deltaDir) % 4
			newRow := row + dirs[newDir].Fst
			newCol := col + dirs[newDir].Snd
			newDist := curDist + 1
			if deltaDir != 0 {
				newDist += 1000
			}
			if grid[newRow][newCol] == '.' {
				if distances[newRow][newCol][newDir] == newDist {
					fromStates[newRow][newCol][newDir] = append(fromStates[newRow][newCol][newDir], curItem.Value)
				}
				if distances[newRow][newCol][newDir] > newDist {
					distances[newRow][newCol][newDir] = newDist
					fromStates[newRow][newCol][newDir] = []Triple{curItem.Value}
					heap.Push(&pq, &Item{
						Value:    types.NewIntTriple(newRow, newCol, newDir),
						Priority: newDist,
					})
				}
			}
		}
	}

	var paths map[Pair]bool = make(map[Pair]bool)
	var toVisit []Triple

	part1Answer = math.MaxInt32
	for i := range 4 {
		if distances[end.Fst][end.Snd][i] == part1Answer {
			toVisit = append(toVisit, types.NewIntTriple(end.Fst, end.Snd, i))
		}
		if distances[end.Fst][end.Snd][i] < part1Answer {
			part1Answer = distances[end.Fst][end.Snd][i]
			toVisit = []Triple{
				types.NewIntTriple(end.Fst, end.Snd, i),
			}
		}
	}

	for len(toVisit) != 0 {
		curState := toVisit[0]
		paths[types.NewIntPair(curState.Fst, curState.Snd)] = true
		toVisit = toVisit[1:]
		toVisit = append(toVisit, fromStates[curState.Fst][curState.Snd][curState.Thd]...)
	}
	part2Answer = len(paths)

	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
}
