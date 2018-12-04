package day13

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day13", func (path string) error {
		lines := input.ReadAsStringSlice(path)
		firewall := NewFirewall()
		for _, line := range lines {
			var d, r int
			_, err := fmt.Sscanf(line, "%d: %d", &d, &r)
			if err != nil {
				return err
			}
			firewall.AddLayer(d, r)
		}
		fmt.Println("Part 1:", firewall.Traverse(0))
		for delay := 0; ; delay++ {
			if firewall.Traverse(delay) == 0 {
				fmt.Println("Part 2:", delay)
				break
			}
		}
		return nil
	})
}
