package day01

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	p := adventofcode.NewPuzzle("day01", func(path string) error {
		changes := input.ReadAsIntSlice(path)
		fmt.Println("Part 1:", part1(changes))
		fmt.Println("Part 2:", part2(changes))
		return nil
	})

	return p
}
