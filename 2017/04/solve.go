package day04

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day04", func (path string) error {
		phrases := input.ReadAsStringSlice(path)
		fmt.Println("Part 1:", partOne(phrases))
		fmt.Println("Part 2:", partTwo(phrases))
		return nil
	})
}

func partOne(phrases []string) (valid int) {
	for _, phrase := range phrases {
		if IsUnique(phrase) {
			valid++
		}
	}

	return valid
}

func partTwo(phrases []string) (valid int) {
	for _, phrase := range phrases {
		if IsNotAnagram(phrase) {
			valid++
		}
	}

	return valid
}
