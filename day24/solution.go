package day24

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/aoc-2024-go/io"
	"github.com/aoc-2024-go/types"
	"github.com/aoc-2024-go/utils"
)

var part1Answer int
var part2Answer string
var sections [][]string

type Operation = types.Triple[string, string, string]
type StrPair = types.Pair[string, string]
type Solution struct{}

var values map[string]int8 = make(map[string]int8)
var operations map[string]Operation = make(map[string]Operation)
var nBits int
var maxZ int
var andOps map[string]StrPair = make(map[string]StrPair)
var orOps map[string]StrPair = make(map[string]StrPair)
var xorOps map[string]StrPair = make(map[string]StrPair)

func (s Solution) Solve(isSample bool, dirName string) (string, string) {
	var pathToInput string

	if isSample {
		pathToInput = fmt.Sprintf("./%s/sample.txt", dirName)
	} else {
		pathToInput = fmt.Sprintf("./%s/input.txt", dirName)
	}
	sections = io.SplitIntoSections(pathToInput)

	inputs := sections[0]
	nBits = len(inputs) / 2
	for _, line := range inputs {
		var inputName string
		var valInt int
		parts := strings.Split(line, ": ")
		inputName = parts[0]
		valInt, _ = strconv.Atoi(parts[1])
		values[inputName] = int8(valInt)
	}

	for _, line := range sections[1] {

		var op1, op, op2, result string
		fmt.Sscanf(line, "%s %s %s -> %s", &op1, &op, &op2, &result)
		if result[0] == 'z' {
			zId, _ := strconv.Atoi(result[1:])
			maxZ = utils.Max(zId, maxZ)
		}

		switch op {
		case "AND":
			andOps[op1] = StrPair{Fst: op2, Snd: result}
			andOps[op2] = StrPair{Fst: op1, Snd: result}
		case "OR":
			orOps[op1] = StrPair{Fst: op2, Snd: result}
			orOps[op2] = StrPair{Fst: op1, Snd: result}

		case "XOR":
			xorOps[op1] = StrPair{Fst: op2, Snd: result}
			xorOps[op2] = StrPair{Fst: op1, Snd: result}
		}
		operations[result] = Operation{Fst: op1, Snd: op, Thd: op2}
	}

	for i := range maxZ + 1 {
		wire := fmt.Sprintf("z%02d", i)
		val := evaluateVar(wire)
		part1Answer = part1Answer | (int(val) << i)
	}

	if isSample {
		return strconv.Itoa(part1Answer), part2Answer
	}

	var anomalies map[string]bool = make(map[string]bool)

	for from, pair := range xorOps {
		if from[0] == 'x' || from[0] == 'y' {
			internal := pair.Snd
			if internal == "z00" {
				continue
			}
			if xorOps[internal] == (StrPair{}) || andOps[internal] == (StrPair{}) {
				anomalies[internal] = true
			}
		} else {
			output := pair.Snd
			if output[0] != 'z' {
				anomalies[output] = true
			}
		}
	}

	for _, pair := range andOps {
		output := pair.Snd
		if output[0] == 'z' {
			anomalies[output] = true
		}
	}

	for _, pair := range orOps {
		output := pair.Snd
		if output[0] == 'z' && output != fmt.Sprintf("z%02d", nBits) {
			anomalies[output] = true
		}
	}

	var sortedAnomalies []string
	for key := range anomalies {
		sortedAnomalies = append(sortedAnomalies, key)
	}
	slices.Sort(sortedAnomalies)
	part2Answer = strings.Join(sortedAnomalies, ",")
	return strconv.Itoa(part1Answer), part2Answer
}

func evaluateVar(wire string) int8 {
	_, evaluated := values[wire]
	if !evaluated {
		var operation Operation = operations[wire]
		op1, op, op2 := operation.Fst, operation.Snd, operation.Thd
		switch op {
		case "AND":
			values[wire] = evaluateVar(op1) & evaluateVar(op2)
		case "OR":
			values[wire] = evaluateVar(op1) | evaluateVar(op2)
		case "XOR":
			values[wire] = evaluateVar(op1) ^ evaluateVar(op2)
		}
	}
	return values[wire]
}
