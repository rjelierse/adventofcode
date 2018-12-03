package day01

import (
	"github.com/rjelierse/adventofcode/grid"
	"strconv"
)

func part2(steps [][]byte) int {
	visits := make(map[grid.Position]bool)
	visits[grid.Position{0, 0}] = true
	iterator := grid.NewIterator()
	for _, step := range steps {
		switch step[0] {
		case 'L':
			iterator.TurnLeft()
		case 'R':
			iterator.TurnRight()
		}
		distance, _ := strconv.Atoi(string(step[1:]))
		for i := 0; i < distance; i++ {
			iterator.MoveOne()
			if visits[iterator.Position] {
				return iterator.DistanceFromCentre()
			}
			visits[iterator.Position] = true
		}
	}
	return -1
}
