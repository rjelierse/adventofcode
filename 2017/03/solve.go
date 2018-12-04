package day03

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day03", func (path string) error {
		value := input.ReadAsInt(path)
		x, y := position(value)
		var carry int

		if x >= 0 && y >= 0 {
			carry = x + y
		} else if x >= 0 && y < 0 {
			carry = x + -1*y
		} else if x < 0 && y >= 0 {
			carry = -1*x + y
		} else if x < 0 && y < 0 {
			carry = -1*x + -1*y
		}

		fmt.Printf("Value at position (%d,%d) - carry %d\n", x, y, carry)
		return nil
	})
}
