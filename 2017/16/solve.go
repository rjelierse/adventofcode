package day16

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day16", func (path string) error {
		instructions, err := ParseInstructions(input.ReadAsString(path))
		if err != nil {
			return err
		}

		floor := NewFloor(16)
		history := NewHistory()

		for i := 0; i < 1e9; i++ {
			position := floor.Dance(instructions)

			if i == 0 {
				fmt.Println("Part 1:", position)
			}

			if history.SeenBefore(position) {
				position = history.Get((1e9 % i) - 1)
				fmt.Println("Part 2:", position)
				break
			}

			history.Save(position)
		}

		return nil
	})
}
