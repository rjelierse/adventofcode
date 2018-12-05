package day24

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
	"sort"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day24", func (path string) error {
		lines := input.ReadAsStringSlice(path)
		links, _ := ParseLinks(lines)
		bridges := links.RecursiveSearch(0)
		sort.Sort(bridges)
		sort.Reverse(bridges)
		fmt.Println("Strongest bridge:", bridges[0].Strength())
		return nil
	})
}
