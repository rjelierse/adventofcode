package day04

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func part1(guards GuardMap) int {
	guard := guards.MostAsleep()
	minute, _ := guard.MostSleptMinute()
	return guard.ID * minute
}

func part2(guards GuardMap) int {
	guard := guards.MostSleptAtSameTime()
	minute, _ := guard.MostSleptMinute()
	return guard.ID * minute
}

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day04", func(path string) error {
		lines := input.ReadAsStrings(path)
		entries := NewLogEntrySlice(lines)
		guards := NewGuardMap(entries)
		fmt.Println("Part 1:", part1(guards))
		fmt.Println("Part 2:", part2(guards))
		return nil
	})
}
