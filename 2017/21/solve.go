package day21

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day21", func (path string) error {
		data := input.ReadAsStringSlice(path)
		enh := ParseEnhancements(data)
		img := DefaultImage()
		for i := 0; i < 5; i++ {
			img = img.Enhance(enh)
		}
		fmt.Println("Part 1:", img.Count())
		for i := 5; i < 18; i++ {
			img = img.Enhance(enh)
		}
		fmt.Println("Part 2:", img.Count())
		return nil
	})
}
