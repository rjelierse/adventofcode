package day17

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func part1(steps int) int {
	ring := NewRing()
	ring.Run(steps, 2017)
	return ring.ValueAfter(2017)
}

func part2(steps int) int {
	var index, result int
	for value := 1; value < 50e6; value++ {
		index = (index + steps + 1) % value
		if index == 0 {
			result = value
		}
	}
	return result
}

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day17", func (path string) error {
		steps := input.ReadAsInt(path)
		fmt.Println("Part 1:", part1(steps))
		fmt.Println("Part 2:", part2(steps))
		return nil
	})
}
