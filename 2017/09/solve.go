package day09

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day09", func (path string) error {
		data := input.ReadAsIs(path)
		score, garbage := ReadStream(string(data))
		fmt.Println("Part 1:", score)
		fmt.Println("Part 2:", garbage)
		return nil
	})
}
