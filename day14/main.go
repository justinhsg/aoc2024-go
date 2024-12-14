package main

import (
	"fmt"
	"os"

	"github.com/aoc-2024-go/io"
	"github.com/aoc-2024-go/types"
	"github.com/aoc-2024-go/utils"
)

var part1Answer, part2Answer int

var width, height int
var outputFile string = "out.txt"

func main() {
	args := os.Args[1:]
	var lines []string
	var isSample bool
	var isVerbose bool
	if len(args) == 1 {
		lines = io.ReadFile("./sample.txt")
		width = 11
		height = 7
		isSample = true
	} else {
		lines = io.ReadFile("./input.txt")
		width = 101
		height = 103
		isSample = false
		isVerbose = len(args) != 0
	}

	var points []types.IntPair = make([]types.IntPair, len(lines))
	var velocities []types.IntPair = make([]types.IntPair, len(lines))
	var widthMid int = (width - 1) / 2
	var heightMid int = (height - 1) / 2
	var quads []int = make([]int, 4)
	for i, line := range lines {
		var tpx, tpy, tvx, tvy int
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &tpx, &tpy, &tvx, &tvy)
		points[i] = types.NewIntPair(tpx, tpy)
		velocities[i] = types.NewIntPair(tvx, tvy)
		newX := ((tpx+tvx*100)%width + width) % width
		newY := ((tpy+tvy*100)%height + height) % height
		if newX < widthMid {
			if newY < heightMid {
				quads[0] += 1
			}
			if newY > heightMid {
				quads[1] += 1
			}
		}
		if newX > widthMid {
			if newY < heightMid {
				quads[3] += 1
			}
			if newY > heightMid {
				quads[2] += 1
			}
		}
	}
	part1Answer = quads[0] * quads[1] * quads[2] * quads[3]
	if isSample {
		fmt.Println(part1Answer)
		fmt.Println(part2Answer)
		return
	}

	var grid [][]bool = make([][]bool, height)

	if isVerbose {
		io.PrepareFileForWriting(outputFile)
	}
	for i := 0; i < width*height; i++ {
		if isVerbose {
			io.AppendLine(fmt.Sprintf("Time: %d:", i), outputFile)
		}
		resetGrid(&grid)
		for j := 0; j < len(points); j++ {
			nx, ny := types.AddPair(points[j], velocities[j].MultScalar(i)).Destruct()
			nx = (nx%width + width) % width
			ny = (ny%height + height) % height
			grid[ny][nx] = true
		}
		if isVerbose {
			printGrid(&grid)
		}
		if hasTree(&grid) {
			part2Answer = i
			break
		}

	}

	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
}

func resetGrid(grid *[][]bool) {
	for i := 0; i < height; i++ {

		(*grid)[i] = make([]bool, width)

	}
}

func printGrid(grid *[][]bool) {
	byteArr := make([][]byte, height)

	for i := 0; i < height; i++ {
		byteArr[i] = make([]byte, width)
		for j := 0; j < width; j++ {
			if (*grid)[i][j] {
				byteArr[i][j] = '#'
			} else {
				byteArr[i][j] = '.'
			}
		}
	}
	lines := utils.Map(byteArr, func(row []byte) string {
		return string(row)
	})
	io.AppendLines(lines, outputFile)
}

func hasTree(grid *[][]bool) bool {
	var maxConsec int = 0
	var curConsec int = 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if (*grid)[i][j] {
				curConsec += 1
			} else {
				if curConsec > maxConsec {
					maxConsec = curConsec
				}
				curConsec = 0
			}
		}
		if maxConsec > 10 {
			return true
		}
	}
	return false
}
