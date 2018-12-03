package day05

import "testing"

func TestCalculateSteps(t *testing.T) {
	moves := []int{0, 3, 0, 1, -3}
	steps := part1(moves)

	if steps != 5 {
		t.Errorf("Expected exit in 5 steps, got %d instead", steps)
	}
}

func TestCalculateSteps2(t *testing.T) {
	moves := []int{0, 3, 0, 1, -3}
	steps := part2(moves)

	if steps != 10 {
		t.Errorf("Expected exit in 10 steps, got %d instead", steps)
	}
}
