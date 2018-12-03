package day02

import (
	"github.com/rjelierse/adventofcode/grid"
)

type Keypad struct {
	grid.Iterator
	buttons [][]byte
}

func NewSquareKeypad() *Keypad {
	return &Keypad{
		Iterator: *grid.NewIteratorAtPosition(1, 1),
		buttons: [][]byte{
			{'1', '2', '3'},
			{'4', '5', '6'},
			{'7', '8', '9'},
		},
	}
}

func NewDiamondKeypad() *Keypad {
	return &Keypad{
		Iterator: *grid.NewIteratorAtPosition(0, 2),
		buttons: [][]byte{
			{'0', '0', '1', '0', '0'},
			{'0', '2', '3', '4', '0'},
			{'5', '6', '7', '8', '9'},
			{'0', 'A', 'B', 'C', '0'},
			{'0', '0', 'D', '0', '0'},
		},
	}
}


func (k *Keypad) TryMove(d grid.Direction) {
	maxX, maxY := len(k.buttons[0]) - 1, len(k.buttons) - 1
	p := k.Position.MoveOne(d)
	if !p.InBounds(0, maxX, 0, maxY) {
		return
	}
	x, y := p.Coordinates()
	if k.buttons[y][x] == '0' {
		return
	}
	k.Position = p
}

func (k *Keypad) GetKey() byte {
	x, y := k.Current()
	return k.buttons[y][x]
}
