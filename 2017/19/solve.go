package day19

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day19", func (path string) error {
		data := input.ReadAsSplitLines(path)
		maze := NewMaze(data)
		sequence, steps := maze.Travel()
		fmt.Println("Part 1:", string(sequence))
		fmt.Println("Part 2:", steps)
		return nil
	})
}
