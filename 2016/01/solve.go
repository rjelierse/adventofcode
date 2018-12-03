package day01

import (
	"bytes"
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day01", func(path string) error {
	    steps := bytes.Split(input.ReadAsIs(path), []byte(", "))
	    fmt.Println("Part 1:", part1(steps))
	    fmt.Println("Part 2:", part2(steps))
		return nil
	})
}
