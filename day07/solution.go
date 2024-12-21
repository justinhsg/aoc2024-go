package day07

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/aoc-2024-go/io"
	"github.com/aoc-2024-go/utils"
)

var part1Answer, part2Answer int

type Solution struct{}

func (s Solution) Solve(isSample bool, dirName string) (string, string) {
	var pathToInput string

	if isSample {
		pathToInput = fmt.Sprintf("./%s/sample.txt", dirName)
	} else {
		pathToInput = fmt.Sprintf("./%s/input.txt", dirName)
	}
	lines := io.ReadFile(pathToInput)

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		target, _ := strconv.Atoi(parts[0])
		rawNumbers := strings.Split(parts[1], " ")
		numbers := utils.Map(rawNumbers, func(raw string) int {
			num, _ := strconv.Atoi(raw)
			return num
		})
		if tryReduce(target, numbers) {
			part1Answer += target
			part2Answer += target
		} else if tryReducePart2(target, numbers) {
			part2Answer += target
		}

	}

	return strconv.Itoa(part1Answer), strconv.Itoa(part2Answer)
}

func tryReduce(target int, slc []int) bool {
	if len(slc) == 1 {
		return target == slc[0]
	}

	lastIndex := len(slc) - 1
	lastNum := slc[lastIndex]

	return (target%lastNum == 0 && tryReduce(target/lastNum, slc[:lastIndex])) ||
		(target > lastNum && tryReduce(target-lastNum, slc[:lastIndex]))

}

func tryReducePart2(target int, slc []int) bool {
	if len(slc) == 1 {
		return target == slc[0]
	}

	lastIndex := len(slc) - 1
	lastNum := slc[lastIndex]
	tenPow := int(math.Pow10(utils.NDigits(lastNum)))
	isConcat := target%tenPow == lastNum

	isValid := (target%lastNum == 0 && tryReducePart2(target/lastNum, slc[:lastIndex])) ||
		(target > lastNum && tryReducePart2(target-lastNum, slc[:lastIndex])) ||
		(isConcat && tryReducePart2(target/tenPow, slc[:lastIndex]))

	return isValid
}
