package day07

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day07", func(path string) error {
		addresses := input.ReadAsSplitLines(path)
		var part1, part2 int
		for _, address := range addresses {
			if SupportsTLS(address) {
				part1++
			}
			if SupportsSSL(address) {
				part2++
			}
		}
		fmt.Println("Part 1:", part1)
		fmt.Println("Part 2:", part2)
		return nil
	})
}
