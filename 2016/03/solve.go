package day03

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func parseInput(lines []string) (values [][]int) {
	values = make([][]int, len(lines))
	for i, line := range lines {
		var a, b, c int
		if _, err := fmt.Sscanln(line, &a, &b, &c); err != nil {
			panic(err)
		}
		values[i] = []int{a, b, c}
	}
	return
}

func part1(values [][]int) (possible int) {
	for _, line := range values {
		a, b, c := line[0], line[1], line[2]
		if isTriangle(a, b, c) {
			possible++
		}
	}
	return
}

func part2(values [][]int) (possible int) {
	for x := 0; x < 3; x++ {
		for y := 0; y < len(values); y += 3 {
			a, b, c := values[y][x], values[y+1][x], values[y+2][x]
			if isTriangle(a, b, c) {
				possible++
			}
		}
	}
	return
}

func isTriangle(a, b, c int) bool {
	return a+b>c && a+c>b && b+c>a
}

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day03", func(path string) error {
		lines := input.ReadAsStringSlice(path)
		values := parseInput(lines)
	    fmt.Println("Part 1:", part1(values))
	    fmt.Println("Part 2:", part2(values))
		return nil
	})
}
