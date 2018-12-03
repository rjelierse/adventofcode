package day10

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day10", func (path string) error {
		data := input.ReadAsIs(path)
		lengths := Length(data)
		for _, b := range Hash(lengths) {
			fmt.Printf("%02x", b)
		}
		fmt.Println()
		return nil
	})
}
