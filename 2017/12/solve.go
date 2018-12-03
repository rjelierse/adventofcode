package day12

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day12", func (path string) error {
		lines := input.ReadAsStrings(path)
		tree, err := ParseTree(lines)
		if err != nil {
			return err
		}
		fmt.Println("Part 1:", tree.Size(0))
		fmt.Println("Part 2:", tree.NumGroups())
		return nil
	})
}
