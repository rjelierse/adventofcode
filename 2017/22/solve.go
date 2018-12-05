package day22

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

func part1(input [][]byte, rounds int) int {
	m := NewMap(input)
	v := NewVirus()

	for i := 0; i < rounds; i++ {
		switch m.GetState(v.Position) {
		case StateInfected:
			v.TurnRight()
			m.SetState(v.Position, StateClean)
		case StateClean:
			v.TurnLeft()
			m.SetState(v.Position, StateInfected)
			v.AddInfection()
		default:
			panic("unknown state")
		}
		v.MoveOne()
	}
	return v.Infections
}

func part2(input [][]byte, rounds int) int {
	m := NewMap(input)
	v := NewVirus()

	for i := 0; i < rounds; i++ {
		switch m.GetState(v.Position) {
		case StateInfected:
			v.TurnRight()
			m.SetState(v.Position, StateFlagged)
		case StateClean:
			v.TurnLeft()
			m.SetState(v.Position, StateWeakened)
		case StateWeakened:
			m.SetState(v.Position, StateInfected)
			v.AddInfection()
		case StateFlagged:
			v.Reverse()
			m.SetState(v.Position, StateClean)
		default:
			panic("unknown state")
		}
		v.MoveOne()
	}
	return v.Infections
}

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day22", func (path string) error {
		lines := input.ReadAsSplitLines(path)
		fmt.Println("[part1] Infections:", part1(lines, 1e4))
		fmt.Println("[part2] Infections:", part2(lines, 1e7))
		return nil
	})
}
