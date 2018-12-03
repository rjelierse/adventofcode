package day01

import (
	"github.com/rjelierse/adventofcode/grid"
	"strconv"
)

func part1(steps [][]byte) int {
	iterator := grid.NewIterator()
	for _, step := range steps {
		switch step[0] {
		case 'L':
			iterator.TurnLeft()
		case 'R':
			iterator.TurnRight()
		}
		distance, _ := strconv.Atoi(string(step[1:]))
		iterator.Move(distance)
	}
	return iterator.DistanceFromCentre()
}
