package day05

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day05", func (path string) error {
		moves := input.ReadAsInts(path)
		fmt.Println("Part 1:", part1(moves))
		fmt.Println("Part 2:", part2(moves))
		return nil
	})
}
