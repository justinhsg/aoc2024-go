package gridutils

import (
	"github.com/aoc-2024-go/types"
)

type Pair = types.IntPair

func GenerateDirs(dist int) []Pair {
	dirs := make([]Pair, 4*dist)
	row := -dist
	col := 0
	acc := 0
	for row != 0 {
		dirs[acc] = Pair{Fst: row, Snd: col}
		acc += 1
		row += 1
		col += 1
	}
	for col != 0 {
		dirs[acc] = Pair{Fst: row, Snd: col}
		acc += 1
		row += 1
		col -= 1
	}
	for row != 0 {
		dirs[acc] = Pair{Fst: row, Snd: col}
		acc += 1
		row -= 1
		col -= 1
	}
	for col != 0 {
		dirs[acc] = Pair{Fst: row, Snd: col}
		acc += 1
		row -= 1
		col += 1
	}
	return dirs
}
