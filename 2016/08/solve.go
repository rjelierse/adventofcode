package day08

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
	"strings"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day08", func(path string) error {
		instructions := input.ReadAsStringSlice(path)
		screen := NewScreen(50, 6)
		for _, instruction := range instructions {
			fields := strings.Fields(instruction)
			switch fields[0] {
			case "rect":
				NewFillInstruction(instruction).Apply(screen)
			case "rotate":
				switch fields[1] {
				case "row":
					NewRotateRowInstruction(instruction).Apply(screen)
				case "column":
					NewRotateColumnInstruction(instruction).Apply(screen)
				}
			}
		}
		fmt.Println("Part 1:", screen.Count())
		screen.Print()
		return nil
	})
}
