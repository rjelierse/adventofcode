package day07

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day07", func (path string) error {
		lines := input.ReadAsStringSlice(path)
		node := BuildTree(lines)
		fmt.Printf("Root: %s\n", node.Name)

		if _, err := node.Weight(); err != nil {
			fmt.Printf("Imbalanced node: %v\n", err)
		}
		return nil
	})
}
