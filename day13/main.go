package main

import (
	"fmt"
	"os"

	"github.com/aoc-2024-go/io"
)

var part1Answer, part2Answer int

func main() {
	args := os.Args[1:]
	var sections [][]string

	if len(args) > 0 {
		sections = io.SplitIntoSections("./sample.txt")
	} else {
		sections = io.SplitIntoSections("./input.txt")
	}

	for _, section := range sections {
		var ax, ay, bx, by, cx, cy int

		fmt.Sscanf(section[0], "Button A: X+%d, Y+%d", &ax, &ay)
		fmt.Sscanf(section[1], "Button B: X+%d, Y+%d", &bx, &by)
		fmt.Sscanf(section[2], "Prize: X=%d, Y=%d", &cx, &cy)

		a, b := solve(ax, ay, bx, by, cx, cy)
		if a >= 0 && b >= 0 {
			part1Answer += a*3 + b
		}

		a2, b2 := solve(ax, ay, bx, by, cx+10000000000000, cy+10000000000000)
		if a2 >= 0 && b2 >= 0 {
			part2Answer += a2*3 + b2
		}
	}

	fmt.Println(part1Answer)
	fmt.Println(part2Answer)
}

// Solves the linear system:
// a1 * a + b1 * b = c2
// a2 * a + b2 * b = c2
func solve(a1 int, a2 int, b1 int, b2 int, c1 int, c2 int) (int, int) {
	// Solve using Kramers Rule: for Mx = c, where x = (a, b)
	//(a = det(Ma) / det(M), b = det(Mb) / det(M))
	det := a1*b2 - a2*b1
	if det == 0 {
		// the equations are linearly independent, luckily none encountered in the given input
		return -1, -1
	}
	// determinant of the matrix by replacing ai with ci
	detA := c1*b2 - c2*b1
	// determinant of the matrix by replacing bi with ci
	detB := a1*c2 - a2*c1

	if detA%det == 0 && detB%det == 0 {
		// There is one integer solution
		return detA / det, detB / det
	} else {
		// There are no solutions
		return -2, -2
	}
}
