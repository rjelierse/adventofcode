package day08

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day08", func (path string) error {
		lines := input.ReadAsStringSlice(path)
		table := NewTable()
		var value int
		for _, line := range lines {
			instruction, condition := ParseLine(line)
			if table.Matches(condition) {
				table.Apply(instruction)
			}
			v := table.LargestValue()
			if v > value {
				value = v
			}
		}
		fmt.Println("Part 1:", value)
		fmt.Println("Part 2:", table.LargestValue())
		return nil
	})
}
