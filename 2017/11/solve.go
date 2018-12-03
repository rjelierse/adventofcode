package day11

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day11", func (path string) error {
		moves := input.ReadAsStrings(path)
		var dist, max int
		g := NewGrid()
		for _, move := range moves {
			g.Move(move)
			dist = g.Distance()

			if dist > max {
				max = dist
			}
		}
		fmt.Println("Part 1:", max)
		fmt.Println("Part 2:", dist)
		return nil
	})
}
