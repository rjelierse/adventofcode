package day05

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func part1(polymer []byte) int {
	return NewPolymer(polymer).Fold().Len()
}

func part2(polymer []byte) (m int) {
	counts := map[byte]int{}
	var char byte
	for char = 'A'; char <= 'Z'; char++ {
		counts[char] = NewPolymer(polymer).Filter(char).Fold().Len()
	}
	for _, l := range counts {
		if m == 0 || l < m {
			m = l
		}
	}
	return
}

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day05", func(path string) error {
		polymer := input.ReadAsIs(path)
		fmt.Println("Part 1:", part1(polymer))
		fmt.Println("Part 2:", part2(polymer))
		return nil
	})
}
