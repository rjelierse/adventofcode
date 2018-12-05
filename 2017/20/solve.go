package day20

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
	"sort"
)

func part1(input []string) int {
	buffer, err := BufferFromInput(input)
	if err != nil {
		panic(err)
	}

	return buffer.Closest(500).Id
}

func part2(data []string) int {
	buffer, err := BufferFromInput(data)
	if err != nil {
		panic(err)
	}

	sort.Slice(buffer.particles, func (i, j int) bool {
		a, b := buffer.particles[i], buffer.particles[j]

		switch {
		case a.Acceleration.Sum() < b.Acceleration.Sum():
			return true
		case a.Velocity.Sum() < b.Velocity.Sum():
			return true
		case a.Position.Sum() < b.Position.Sum():
			return true
		default:
			return false
		}
	})

	for t := 0; buffer.Run(t); t++ {
	}

	return buffer.Count(false)
}

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day20", func (path string) error {
		particles := input.ReadAsStringSlice(path)
		fmt.Println("Part 1:", part1(particles))
		fmt.Println("Part 2:", part2(particles))
		return nil
	})
}
