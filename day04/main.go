package main

import (
	"fmt"
	"os"

	"github.com/aoc-2024-go/io"
)

func main() {
	args := os.Args[1:]
	var lines []string

	if len(args) > 0 {
		lines = io.ReadFile("./sample.txt")
	} else {
		lines = io.ReadFile("./input.txt")
	}

	part1Answer := 0
	part2Answer := 0

	dc := []int{0, 1, 1, 1, 0, -1, -1, -1}
	dr := []int{-1, -1, 0, 1, 1, 1, 0, -1}
	xmas := []rune("XMAS")

	width := len(lines[0])
	height := len(lines)

	grid := make([][]rune, height)
	for cRow, line := range lines {
		grid[cRow] = make([]rune, width)
		for cCol, char := range line {
			grid[cRow][cCol] = char
		}
	}

	for cRow := 0; cRow < height; cRow++ {
		for cCol := 0; cCol < width; cCol++ {
			if grid[cRow][cCol] == xmas[0] {
				for i := 0; i < 8; i++ {
					isXmas := true
					for j := 1; j < 4; j++ {
						nRow := cRow + dr[i]*j
						nCol := cCol + dc[i]*j

						if nRow >= 0 && nRow < height && nCol >= 0 && nCol < width {
							tRune := grid[nRow][nCol]
							if tRune != xmas[j] {
								isXmas = false
								break
							}
						} else {
							isXmas = false
						}
					}
					if isXmas {
						part1Answer += 1
					}
				}

			}
		}
	}

	for cRow := 1; cRow < height-1; cRow++ {
		for cCol := 1; cCol < width-1; cCol++ {
			if grid[cRow][cCol] == xmas[2] {

				tl := grid[cRow+dr[7]][cCol+dc[7]]
				tr := grid[cRow+dr[1]][cCol+dc[1]]
				br := grid[cRow+dr[3]][cCol+dc[3]]
				bl := grid[cRow+dr[5]][cCol+dc[5]]

				diag1IsMas := (tl == xmas[1] && br == xmas[3]) ||
					(tl == xmas[3] && br == xmas[1])
				diag2IsMas := (tr == xmas[1] && bl == xmas[3]) ||
					(tr == xmas[3] && bl == xmas[1])

				if diag1IsMas && diag2IsMas {
					part2Answer += 1
				}

			}
		}
	}

	fmt.Println(part1Answer)
	fmt.Println(part2Answer)

}
