package day14

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day14", func (path string) error {
		grid := NewGrid(input.ReadAsString(path))
		fmt.Println("Part 1:", grid.Count())
		fmt.Println("Part 2:", grid.GroupCount())
		return nil
	})
}
