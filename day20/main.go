package day20

import (
	"fmt"
	"math"
	"strconv"

	"github.com/aoc-2024-go/gridutils"
	"github.com/aoc-2024-go/io"
	"github.com/aoc-2024-go/types"
)

var part1Answer, part2Answer int

type Pair = types.IntPair
type State = types.Pair[Pair, int]

var grid [][]byte
var distancesToEnd [][]int
var distancesToStart [][]int
var width, height int
var maxCheat int = 20
var dirs [][]Pair = make([][]types.IntPair, maxCheat+1)

var start, end Pair

type Solution struct{}

func (s Solution) Solve(isSample bool, dirName string) (string, string) {
	var pathToInput string

	if isSample {
		pathToInput = fmt.Sprintf("./%s/sample.txt", dirName)
	} else {
		pathToInput = fmt.Sprintf("./%s/input.txt", dirName)
	}
	lines := io.ReadFile(pathToInput)
	for i := 1; i <= maxCheat; i++ {
		dirs[i] = gridutils.GenerateDirs(i)
	}
	height, width = len(lines), len(lines[0])
	grid = make([][]byte, height)
	distancesToEnd = make([][]int, height)
	distancesToStart = make([][]int, height)

	for row, line := range lines {
		grid[row] = make([]byte, width)
		distancesToEnd[row] = make([]int, width)
		distancesToStart[row] = make([]int, width)
		for col := range width {
			grid[row][col] = line[col]
			distancesToEnd[row][col] = math.MaxInt64
			distancesToStart[row][col] = math.MaxInt64
			if line[col] == 'E' {
				end = Pair{Fst: row, Snd: col}
			}
			if line[col] == 'S' {
				start = Pair{Fst: row, Snd: col}
			}
		}
	}
	fillDistances(start, &distancesToStart)
	fillDistances(end, &distancesToEnd)
	usualTime := distancesToEnd[start.Fst][start.Snd]
	for row := range height {
		for col := range width {
			if grid[row][col] == '#' {
				continue
			}
			startCheat := Pair{Fst: row, Snd: col}
			startRow, startCol := startCheat.Destruct()

			for cheatDuration := 2; cheatDuration < 21; cheatDuration++ {
				for _, dir := range dirs[cheatDuration] {
					endRow, endCol := types.AddPair(startCheat, dir).Destruct()
					if endRow >= 0 && endRow < height && endCol >= 0 && endCol < width && grid[endRow][endCol] != '#' {
						cheatTime := distancesToEnd[endRow][endCol] + distancesToStart[startRow][startCol] + cheatDuration
						if usualTime-cheatTime >= 100 {
							part2Answer += 1
							if cheatDuration == 2 {
								part1Answer += 1
							}
						}
					}
				}
			}
		}
	}
	return strconv.Itoa(part1Answer), strconv.Itoa(part2Answer)
}

func fillDistances(start Pair, distanceGrid *[][]int) {
	(*distanceGrid)[start.Fst][start.Snd] = 0
	var toVisit []State = []State{{Fst: start, Snd: 0}}
	for len(toVisit) != 0 {
		curState := toVisit[0]
		toVisit = toVisit[1:]
		curPos, curDist := curState.Destruct()
		for _, dir := range dirs[1] {
			newPos := types.AddPair(curPos, dir)
			newRow, newCol := newPos.Destruct()
			if newRow >= 0 && newRow < height && newCol >= 0 && newCol < width {
				if grid[newRow][newCol] != '#' && (*distanceGrid)[newRow][newCol] > curDist+1 {
					(*distanceGrid)[newRow][newCol] = curDist + 1
					toVisit = append(toVisit, State{Fst: newPos, Snd: curDist + 1})
				}
			}
		}

	}
}
