package day02

import "github.com/rjelierse/adventofcode/grid"

func part2(moves [][]byte) string {
	var pin []byte
	pad := NewDiamondKeypad()
	for _, line := range moves {
		for _, move := range line {
			switch move {
			case 'L':
				pad.TryMove(grid.DirectionLeft)
			case 'R':
				pad.TryMove(grid.DirectionRight)
			case 'U':
				pad.TryMove(grid.DirectionUp)
			case 'D':
				pad.TryMove(grid.DirectionDown)
			}
		}
		pin = append(pin, pad.GetKey())
	}
	return string(pin)
}
