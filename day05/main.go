package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/aoc-2024-go/io"
)

func main() {
	args := os.Args[1:]
	var sections [][]string

	if len(args) > 0 {
		sections = io.SplitIntoSections("./sample.txt")
	} else {
		sections = io.SplitIntoSections("./input.txt")
	}

	nodes := make(map[string]bool)
	adjMat := make(map[string]map[string]int)
	for _, edge := range sections[0] {
		parts := strings.Split(edge, "|")
		nodes[parts[0]] = true
		nodes[parts[1]] = true
	}

	for i := range maps.Keys(nodes) {
		adjMat[i] = make(map[string]int)
	}

	for _, edge := range sections[0] {
		parts := strings.Split(edge, "|")
		adjMat[parts[0]][parts[1]] = 1
	}

	part1Answer := 0
	part2Answer := 0

	for _, line := range sections[1] {
		proposedPath := strings.Split(line, ",")
		isPathValid := true
		for i := 0; i < len(proposedPath)-1; i++ {
			from := proposedPath[i]
			to := proposedPath[i+1]
			if adjMat[from][to] != 1 {
				isPathValid = false
				break
			}
		}
		if isPathValid {
			midpt := (len(proposedPath) - 1) / 2
			midVal, _ := strconv.Atoi(proposedPath[midpt])
			part1Answer += midVal
		} else {
			slices.SortFunc(proposedPath, func(a, b string) int {
				if adjMat[a][b] == 1 {
					return -1
				} else {
					return 1
				}
			})
			midpt := (len(proposedPath) - 1) / 2
			midVal, _ := strconv.Atoi(proposedPath[midpt])
			part2Answer += midVal
		}
	}

	fmt.Println(part1Answer)
	fmt.Println(part2Answer)

}
