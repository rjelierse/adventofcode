package day02

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day02", func(path string) error {
		boxes := input.ReadAsStrings(path)
		fmt.Println("Part 1:", part1(boxes))
		fmt.Println("Part 2:", part2(boxes))
		return nil
	})
}
