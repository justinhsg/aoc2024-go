package main

import (
	"fmt"
	"os"
	"sort"

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

	var (
		list1 []int
		list2 []int
		t1    int
		t2    int
	)

	list2Counts := make(map[int]int)

	for _, line := range lines {

		fmt.Sscanf(line, "%d %d", &t1, &t2)
		list1 = append(list1, t1)
		list2 = append(list2, t2)
		list2Counts[t2] += 1
	}
	sort.Slice(list1, func(i, j int) bool {
		return list1[i] < list1[j]
	})
	sort.Slice(list2, func(i, j int) bool {
		return list2[i] < list2[j]
	})

	p1Answer := 0
	p2Answer := 0
	for i := 0; i < len(list1); i++ {
		if list1[i] < list2[i] {
			p1Answer += list2[i] - list1[i]
		} else {
			p1Answer += list1[i] - list2[i]
		}
		p2Answer += list1[i] * list2Counts[list1[i]]

	}

	fmt.Println(p1Answer)
	fmt.Println(p2Answer)

}
