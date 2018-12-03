package day02

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day02", func(path string) error {
	    data := input.ReadAsSplitLines(path)
	    fmt.Println("Part 1:", part1(data))
	    fmt.Println("Part 2:", part2(data))
		return nil
	})
}
