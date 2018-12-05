package day18

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day18", func (path string) error {
		instructions := input.ReadAsStringSlice(path)
		fmt.Println("Part 1:", Sounds(instructions))
		fmt.Println("Part 2:", Run(instructions))
		return nil
	})
}
