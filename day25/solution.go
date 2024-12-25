package day25

import (
	"fmt"
	"strconv"

	"github.com/aoc-2024-go/io"
)

var part1Answer, part2Answer int
var sections [][]string

type Tuple5 struct {
	a, b, c, d, e int
}

var allKeys map[Tuple5]bool = make(map[Tuple5]bool)
var allLocks map[Tuple5]bool = make(map[Tuple5]bool)

type Solution struct{}

func (s Solution) Solve(isSample bool, dirName string) (string, string) {
	var pathToInput string

	if isSample {
		pathToInput = fmt.Sprintf("./%s/sample.txt", dirName)
	} else {
		pathToInput = fmt.Sprintf("./%s/input.txt", dirName)
	}
	sections = io.SplitIntoSections(pathToInput)

	for _, section := range sections {
		if section[0] == "....." {
			var keys []int = make([]int, 5)
			for pos := range 5 {
				for h := 5; h >= 0; h-- {
					if section[6-h][pos] == '#' {
						keys[pos] = h
						break
					}
				}
			}
			key := Tuple5{a: keys[0], b: keys[1], c: keys[2], d: keys[3], e: keys[4]}
			if allKeys[key] {
				panic("!!")
			}
			allKeys[key] = true
		} else {
			var locks []int = make([]int, 5)
			for pos := range 5 {
				for h := 5; h >= 0; h-- {
					if section[h][pos] == '#' {
						locks[pos] = h
						break
					}
				}
			}
			lock := Tuple5{a: locks[0], b: locks[1], c: locks[2], d: locks[3], e: locks[4]}
			if allLocks[lock] {
				panic("!")
			}
			allLocks[lock] = true
		}
	}

	for key := range allKeys {
		requiredLock := Tuple5{a: 5 - key.a, b: 5 - key.b, c: 5 - key.c, d: 5 - key.d, e: 5 - key.e}
		for lock := range allLocks {
			if lock.a <= requiredLock.a &&
				lock.b <= requiredLock.b &&
				lock.c <= requiredLock.c &&
				lock.d <= requiredLock.d &&
				lock.e <= requiredLock.e {
				part1Answer += 1
			}
		}
		// if allLocks[requiredLock] {
		// 	part1Answer += 1
		// }

	}
	fmt.Println(allKeys, allLocks)
	fmt.Println(len(allKeys) + len(allLocks))

	return strconv.Itoa(part1Answer), "Merry Christmas!"
}
