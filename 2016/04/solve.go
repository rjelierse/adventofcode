package day04

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day04", func(path string) error {
	    lines := input.ReadAsSplitLines(path)
	    var sum int
	    var storage int
	    for _, line := range lines {
	    	room := NewRoom(line)
	    	if !room.IsValid() {
	    		continue
			}
			sum += room.ID
			if room.Name() == "northpole object storage" {
				storage = room.ID
			}
		}
	    fmt.Println("Part 1:", sum)
	    fmt.Println("Part 2:", storage)
		return nil
	})
}
