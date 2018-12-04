package day09

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day09", func(path string) error {
	    data := input.ReadAsIs(path)
	    fmt.Println("Part 1:", part1(data).DecompressedLength())
	    fmt.Println("Part 2:", part2(data).DecompressedLength())
		return nil
	})
}
