package day20

import (
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day20", func(path string) error {
		// TODO: Implement
		return nil
	})
}
