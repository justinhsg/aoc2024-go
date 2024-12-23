package day23

import (
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"

	"github.com/aoc-2024-go/io"
	"github.com/aoc-2024-go/utils"
)

var part1Answer int
var part2Answer string
var lines []string

type Solution struct{}

var nodeToIdx map[string]int = make(map[string]int)
var nodes []string

var nodeMap map[string]bool = make(map[string]bool)
var adjList [][]int
var adjMat map[int]map[int]bool = make(map[int]map[int]bool)

func (s Solution) Solve(isSample bool, dirName string) (string, string) {
	var pathToInput string

	if isSample {
		pathToInput = fmt.Sprintf("./%s/sample.txt", dirName)
	} else {
		pathToInput = fmt.Sprintf("./%s/input.txt", dirName)
	}
	lines = io.ReadFile(pathToInput)

	for _, line := range lines {
		parts := strings.Split(line, "-")
		node1, node2 := parts[0], parts[1]
		nodeMap[node1] = true
		nodeMap[node2] = true
	}
	nodes = make([]string, len(nodeMap))
	adjList = make([][]int, len(nodeMap))
	i := 0
	for node := range maps.Keys(nodeMap) {
		nodes[i] = node
		i += 1
	}
	slices.Sort(nodes)
	for i, n1 := range nodes {
		nodeToIdx[n1] = i
		adjMat[i] = make(map[int]bool)
	}
	for _, line := range lines {
		parts := strings.Split(line, "-")
		n1, n2 := nodeToIdx[parts[0]], nodeToIdx[parts[1]]
		if n1 < n2 {
			adjList[n1] = append(adjList[n1], n2)
			adjMat[n1][n2] = true
		} else {
			adjList[n2] = append(adjList[n2], n1)
			adjMat[n2][n1] = true
		}

	}

	for i := range len(nodes) {
		for _, j := range adjList[i] {
			for _, k := range adjList[j] {
				if adjMat[i][k] {
					if nodes[i][0] == 't' || nodes[j][0] == 't' || nodes[k][0] == 't' {
						part1Answer += 1
					}
				}
			}
		}
	}

	var toConsider [][]int = make([][]int, len(nodes))
	for i := range len(nodes) {
		toConsider[i] = []int{i}
	}

	for {
		nextGen := findConnection(&toConsider)
		if len(nextGen) == 0 {
			nodeArr := utils.Map(toConsider[0], func(x int) string { return nodes[x] })
			part2Answer = strings.Join(nodeArr, ",")
			break
		} else {
			toConsider = nextGen
		}
	}

	return strconv.Itoa(part1Answer), part2Answer
}

func findConnection(toConsider *[][]int) [][]int {
	var nextGen [][]int = [][]int{}
	for _, curSeq := range *toConsider {
		lastIndex := curSeq[len(curSeq)-1]
		for _, candidate := range adjList[lastIndex] {
			if candidate < lastIndex {
				continue
			}
			isConnected := true
			for _, inCluster := range curSeq {
				if !adjMat[inCluster][candidate] {
					isConnected = false
					break
				}
			}
			if isConnected {
				nextSeq := make([]int, len(curSeq)+1)
				_ = copy(nextSeq, curSeq)
				nextSeq[len(curSeq)] = candidate
				nextGen = append(nextGen, nextSeq)
			}
		}
	}
	return nextGen
}
