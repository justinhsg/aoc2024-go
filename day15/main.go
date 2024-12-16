package main

import (
	"fmt"
	"maps"
	"os"

	"github.com/aoc-2024-go/io"
	"github.com/aoc-2024-go/types"
)

var part1Answer, part2Answer int

var width, height int
var outputFile string = "out.txt"

func main() {
	args := os.Args[1:]
	var sections [][]string

	if len(args) == 1 {
		sections = io.SplitIntoSections("./sample.txt")

	} else {
		sections = io.SplitIntoSections("./input.txt")
	}

	width := len(sections[0][0])
	wideWidth := width * 2
	height := len(sections[0])

	grid := make([][]byte, height)
	wideGrid := make([][]byte, height)

	var start, wideStart types.IntPair
	var dirs map[rune]types.IntPair
	dirs = map[rune]types.IntPair{
		'^': types.NewIntPair(-1, 0),
		'>': types.NewIntPair(0, 1),
		'v': types.NewIntPair(1, 0),
		'<': types.NewIntPair(0, -1),
	}
	for row, line := range sections[0] {
		grid[row] = make([]byte, width)
		wideGrid[row] = make([]byte, wideWidth)
		for col, ch := range line {
			grid[row][col] = byte(ch)
			wideGrid[row][2*col] = byte(ch)
			wideGrid[row][2*col+1] = byte(ch)
			if ch == '@' {
				start = types.NewIntPair(row, col)
				wideStart = types.NewIntPair(row, col*2)
				grid[row][col] = '.'
				wideGrid[row][2*col] = '.'
				wideGrid[row][2*col+1] = '.'
			}
			if ch == 'O' {
				wideGrid[row][2*col] = '['
				wideGrid[row][2*col+1] = ']'
			}

		}
	}
	curPos := start
	curPos2 := wideStart
	for _, line := range sections[1] {
		for _, ch := range line {
			curPos = tryMove(curPos, dirs[ch], &grid)
			curPos2 = tryMove2(curPos2, dirs[ch], &wideGrid)
		}
	}

	for row := range height {
		for col := range width {
			if grid[row][col] == 'O' {
				part1Answer += 100*row + col
			}
		}
	}

	for row := range height {
		for col := range wideWidth {
			if wideGrid[row][col] == '[' {
				part2Answer += 100*row + col
			}
		}
	}

	fmt.Println(part1Answer)
	fmt.Println(part2Answer)

}

func tryMove(curPos types.IntPair, dir types.IntPair, grid *[][]byte) types.IntPair {
	newRow, newCol := types.AddPair(curPos, dir).Destruct()
	if (*grid)[newRow][newCol] == '.' {
		return types.NewIntPair(newRow, newCol)
	}
	if (*grid)[newRow][newCol] == '#' {
		return curPos
	}
	rockRow := newRow
	rockCol := newCol
	for (*grid)[rockRow][rockCol] == 'O' {
		rockRow += dir.Fst
		rockCol += dir.Snd
	}
	if (*grid)[rockRow][rockCol] == '.' {
		(*grid)[rockRow][rockCol] = 'O'
		(*grid)[newRow][newCol] = '.'
		return types.NewIntPair(newRow, newCol)
	} else {
		return curPos
	}

}

func tryMove2(curPos types.IntPair, dir types.IntPair, wideGrid *[][]byte) types.IntPair {
	newRow, newCol := types.AddPair(curPos, dir).Destruct()
	if (*wideGrid)[newRow][newCol] == '.' {
		return types.NewIntPair(newRow, newCol)
	}
	if (*wideGrid)[newRow][newCol] == '#' {
		return curPos
	}
	var hasPush bool
	if dir.Fst == 0 {
		hasPush = tryPushHoriz(newRow, newCol, dir.Snd, wideGrid)

	} else {
		hasPush = tryPushVert(newRow, newCol, dir.Fst, wideGrid)
	}
	if hasPush {
		return types.NewIntPair(newRow, newCol)
	} else {
		return curPos
	}
}

func tryPushHoriz(row int, col int, delta int, wideGrid *[][]byte) bool {
	endCol := col
	stillRock := true
	for stillRock {
		endCol += delta
		stillRock = (*wideGrid)[row][endCol] == '[' || (*wideGrid)[row][endCol] == ']'
	}
	if (*wideGrid)[row][endCol] == '#' {
		return false
	}
	for endCol != col {
		(*wideGrid)[row][endCol] = (*wideGrid)[row][endCol-delta]
		(*wideGrid)[row][endCol-delta] = '.'
		endCol -= delta
	}
	return true
}

func tryPushVert(row int, col int, delta int, wideGrid *[][]byte) bool {
	var considered map[int]map[int]bool = make(map[int]map[int]bool)
	var toConsider map[int]bool = make(map[int]bool)

	toConsider[col] = true
	if (*wideGrid)[row][col] == '[' {
		toConsider[col+1] = true
	} else {
		toConsider[col-1] = true
	}
	var isBlocked bool = false

	curRow := row
	for len(toConsider) != 0 && !isBlocked {
		nextRow := curRow + delta
		newToConsider := make(map[int]bool)
		for considerCol := range maps.Keys(toConsider) {
			nextSpace := (*wideGrid)[nextRow][considerCol]
			if nextSpace == '[' {
				newToConsider[considerCol] = true
				newToConsider[considerCol+1] = true
			}
			if nextSpace == ']' {
				newToConsider[considerCol] = true
				newToConsider[considerCol-1] = true
			}
			if nextSpace == '#' {
				isBlocked = true
				break
			}
		}
		considered[curRow] = toConsider
		toConsider = newToConsider
		curRow = nextRow
	}
	if isBlocked {
		return false
	}

	for ; curRow != row; curRow -= delta {
		fromRow := curRow - delta
		for considerCol := range maps.Keys(considered[fromRow]) {
			(*wideGrid)[curRow][considerCol] = (*wideGrid)[fromRow][considerCol]
			(*wideGrid)[fromRow][considerCol] = '.'
		}
	}

	return true
}
