package day03

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day03", func(path string) error {
		var claims []*Claim
		definitions := input.ReadAsStrings(path)
		for _, definition := range definitions {
			claims = append(claims, NewClaim(definition))
		}
		cloth := NewCloth()
		for _, claim := range claims {
			cloth.AddClaim(claim)
		}
		fmt.Println("Part 1:", cloth.CountOverlaps())
		for _, claim := range claims {
			if cloth.IsUnique(claim) {
				fmt.Println("Part 2:", claim.id)
				break
			}
		}
		return nil
	})
}
