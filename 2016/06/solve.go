package day06

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day06", func(path string) error {
	    messages := input.ReadAsSplitLines(path)
	    distribution := NewCharacterDistributionFromMessages(messages)
	    fmt.Printf("Part 1: %s\n", distribution.MostCommon())
	    fmt.Printf("Part 2: %s\n", distribution.LeastCommon())
		return nil
	})
}
