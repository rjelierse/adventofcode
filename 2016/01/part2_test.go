package day01

import (
	"testing"
)

func TestSecondVisit(t *testing.T) {
	steps := [][]byte{
		[]byte("R8"),
		[]byte("R4"),
		[]byte("R4"),
		[]byte("R8"),
	}
	distance := part2(steps)
	if distance != 4 {
		t.Errorf("Expected distance from centre at 4 block, got %d instead", distance)
	}
}
