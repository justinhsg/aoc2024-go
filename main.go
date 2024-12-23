package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/aoc-2024-go/day01"
	"github.com/aoc-2024-go/day02"
	"github.com/aoc-2024-go/day03"
	"github.com/aoc-2024-go/day04"
	"github.com/aoc-2024-go/day05"
	"github.com/aoc-2024-go/day06"
	"github.com/aoc-2024-go/day07"
	"github.com/aoc-2024-go/day08"
	"github.com/aoc-2024-go/day09"
	"github.com/aoc-2024-go/day10"
	"github.com/aoc-2024-go/day11"
	"github.com/aoc-2024-go/day12"
	"github.com/aoc-2024-go/day13"
	"github.com/aoc-2024-go/day14"
	"github.com/aoc-2024-go/day15"
	"github.com/aoc-2024-go/day16"
	"github.com/aoc-2024-go/day17"
	"github.com/aoc-2024-go/day18"
	"github.com/aoc-2024-go/day19"
	"github.com/aoc-2024-go/day20"
	"github.com/aoc-2024-go/day21"
	"github.com/aoc-2024-go/day22"
	"github.com/aoc-2024-go/day23"
	"github.com/aoc-2024-go/day24"
	"github.com/aoc-2024-go/day25"
	"github.com/aoc-2024-go/utils"
	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {
	var day int
	var isSample bool
	var runAll bool
	flag.IntVar(&day, "day", 1, "")
	flag.BoolVar(&isSample, "sample", false, "")
	flag.BoolVar(&runAll, "all", false, "")
	flag.Parse()
	if day < 0 || day > 25 {
		panic("Please provide a day between 1 and 25")
	}

	measurePerformance(day, isSample, runAll)

}

func measurePerformance(day int, isSample bool, runAll bool) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Day", "Part 1", "Part 2", "Duration (ms)"})
	var totalDuration float64
	if !runAll {
		part1, part2, duration := solve(isSample, day)
		t.AppendRow([]interface{}{day, part1, part2, duration})
		totalDuration += duration
	} else {
		for i := range 25 {
			part1, part2, duration := solve(isSample, i+1)
			t.AppendRow([]interface{}{i + 1, part1, part2, duration})
			totalDuration += duration
		}

	}
	t.AppendFooter(table.Row{"", "", "Total", totalDuration})
	t.Render()
}

func solve(isSample bool, day int) (string, string, float64) {
	var solution utils.Solution = getSolution(day)
	var dirName string = fmt.Sprintf("day%02d", day)
	startTime := time.Now()
	part1, part2 := solution.Solve(isSample, dirName)
	endTime := time.Now()
	duration := float64(endTime.Sub(startTime)) / float64(time.Millisecond)
	return part1, part2, duration
}

func getSolution(day int) utils.Solution {
	switch day {
	case 1:
		return day01.Solution{}
	case 2:
		return day02.Solution{}
	case 3:
		return day03.Solution{}
	case 4:
		return day04.Solution{}
	case 5:
		return day05.Solution{}
	case 6:
		return day06.Solution{}
	case 7:
		return day07.Solution{}
	case 8:
		return day08.Solution{}
	case 9:
		return day09.Solution{}
	case 10:
		return day10.Solution{}
	case 11:
		return day11.Solution{}
	case 12:
		return day12.Solution{}
	case 13:
		return day13.Solution{}
	case 14:
		return day14.Solution{}
	case 15:
		return day15.Solution{}
	case 16:
		return day16.Solution{}
	case 17:
		return day17.Solution{}
	case 18:
		return day18.Solution{}
	case 19:
		return day19.Solution{}
	case 20:
		return day20.Solution{}
	case 21:
		return day21.Solution{}
	case 22:
		return day22.Solution{}
	case 23:
		return day23.Solution{}
	case 24:
		return day24.Solution{}
	case 25:
		return day25.Solution{}
	default:
		panic("Day not found")
	}
}
