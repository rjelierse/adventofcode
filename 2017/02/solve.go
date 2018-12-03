package day02

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day03", func (path string) error {
		var sum1, sum2 int
		for _, line := range input.ReadAsStrings(path) {
			numbers := convertLine(line)
			sum1 = sum1 + Difference(numbers)
			sum2 = sum2 + Division(numbers)
		}
		fmt.Println("Part 1:", sum1)
		fmt.Println("Part 2:", sum2)
		return nil
	})
}
