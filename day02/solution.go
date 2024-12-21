package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aoc-2024-go/io"
	"github.com/aoc-2024-go/utils"
)

type Solution struct{}

func isDecreasing(el int) bool {
	return (el >= -3 && el <= -1)
}

func isIncreasing(el int) bool {
	return (el >= 1 && el <= 3)
}

func checkDecreasing(arr []int) bool {
	return utils.All(arr, isDecreasing)
}

func checkIncreasing(arr []int) bool {
	return utils.All(arr, isIncreasing)
}

func mergeBefore(arr []int, idx int) []int {
	newArr := make([]int, len(arr)-1)

	if idx == 0 {
		copy(newArr, arr[1:])
	} else {
		copy(newArr[:idx-1], arr[:idx-1])
		newArr[idx-1] = arr[idx-1] + arr[idx]
		copy(newArr[idx:], arr[idx+1:])
	}
	return newArr
}

func mergeAfter(arr []int, idx int) []int {
	newArr := make([]int, len(arr)-1)

	if idx == len(arr)-1 {
		copy(newArr, arr[:idx-1])
	} else {
		copy(newArr[:idx], arr[:idx])
		newArr[idx] = arr[idx+1] + arr[idx]
		copy(newArr[idx+1:], arr[idx+2:])
	}
	return newArr
}

func checkIncreasingWithTolerance(arr []int) bool {
	idx, _ := utils.Find(arr, func(el int) bool {
		return !isIncreasing(el)
	})
	if idx == -1 || idx == len(arr)-1 {
		return true
	}
	return checkIncreasing(mergeBefore(arr, idx)) || checkIncreasing(mergeAfter(arr, idx))
}

func checkDecreasingWithTolerance(arr []int) bool {
	idx, _ := utils.Find(arr, func(el int) bool {
		return !isDecreasing(el)
	})
	if idx == -1 || idx == len(arr)-1 {
		return true
	}
	return checkDecreasing(mergeBefore(arr, idx)) || checkDecreasing(mergeAfter(arr, idx))
}

func (s Solution) Solve(isSample bool, dirName string) (string, string) {
	var pathToInput string

	if isSample {
		pathToInput = fmt.Sprintf("./%s/sample.txt", dirName)
	} else {
		pathToInput = fmt.Sprintf("./%s/input.txt", dirName)
	}

	lines := io.ReadFile(pathToInput)

	var (
		p1Answer int
		p2Answer int
	)

	for _, line := range lines {
		report := utils.Map(strings.Fields(line), func(el string) int {
			i, _ := strconv.Atoi(el)
			return i
		})
		diffArr := make([]int, len(report)-1)
		for i, value := range report[1:] {
			diffArr[i] = value - report[i]
		}

		if checkIncreasing(diffArr) || checkDecreasing(diffArr) {
			p1Answer += 1
			p2Answer += 1
		} else if checkIncreasingWithTolerance(diffArr) || checkDecreasingWithTolerance(diffArr) {
			p2Answer += 1
		}
	}
	return strconv.Itoa(p1Answer), strconv.Itoa(p2Answer)
}
