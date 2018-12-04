package day05

import (
	"bytes"
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
)

const placeholder = '_'

func part1(prefix []byte) []byte {
	password := make([]byte, 0, 8)
	for i := 0; len(password) < 8; i++ {
		p := NewPassword(prefix, i)
		if !p.IsNext() {
			continue
		}
		password = append(password, p.NextCharacter())
	}
	return password
}

func part2(prefix []byte) (password []byte) {
	password = bytes.Repeat([]byte{placeholder}, 8)
	for i := 0; ; i++ {
		if i % 1e6 == 0 {
			fmt.Printf("Progress: %s\n", password)
		}
		p := NewPassword(prefix, i)
		if !p.IsNext() {
			continue
		}
		pos, c := p.NextCharacterAtPosition()
		if pos < 0 || pos > 7 {
			continue
		}
		if password[pos] == placeholder {
			password[pos] = c
		}
		if !bytes.Contains(password, []byte{placeholder}) {
			return
		}
	}
}

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day05", func(path string) error {
		id := input.ReadAsIs(path)
	    fmt.Printf("Part 1: %s\n", part1(id))
	    fmt.Printf("Part 2: %s\n", part2(id))
		return nil
	})
}
