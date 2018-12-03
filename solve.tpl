package DAY

import (
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("DAY", func(path string) error {
	    // TODO: Implement
		return nil
	})
}
