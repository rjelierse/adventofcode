package day01

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day01", func (path string) error {
		data := input.ReadAsIs(path)
		fmt.Println("Part 1:", calculateSum(data, 1))
		fmt.Println("Part 2:", calculateSum(data, len(data) / 2))
		return nil
	})
}
