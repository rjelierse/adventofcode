package day23

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day23", func (path string) error {
		lines := input.ReadAsStringSlice(path)
		fmt.Println("Part 1:", program(lines))
		fmt.Println("Part 2:", run(79))
		return nil
	})
}

// run is the input assembly written as Go code (with a few optimizations)
func run(b int) int {
	var c, g, h int
	b = (b * 100) + 100000
	c = b + 17000
	for {
		for d := 2; d * d <= b; d++ {
			if b % d == 0 {
				h++
				break
			}
		}
		g = b - c
		b += 17
		if g == 0 {
			break
		}
	}
	return h
}
