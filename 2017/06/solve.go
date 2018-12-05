package day06

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
	"strconv"
	"strings"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day06", func (path string) error {
		slots := strings.Split(input.ReadAsString(path), "\t")
		config := make([]int, len(slots))
		for index, slot := range slots {
			value, err := strconv.Atoi(slot)
			if err != nil {
				panic(err)
			}
			config[index] = value
		}
		var step, first int
		var seen bool

		bank := NewBank(config)
		for step = 0; !seen; step++ {
			bank.Redistribute()
			seen, first = bank.ConfigurationSeenBefore()
		}

		fmt.Println("Steps", step, "length", step-first)
		return nil
	})
}
