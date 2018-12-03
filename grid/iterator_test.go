package grid

import "testing"

type scenario struct {
	p Position
	d int
}

var doubles = []scenario{
	{p: Position{0, 4}, d: 4},
	{p: Position{0, -4}, d: 4},
	{p: Position{-2, -2}, d: 4},
	{p: Position{2, -2}, d: 4},
}

func TestIterator_Current(t *testing.T) {
	i := NewIteratorAtPosition(2, 2)
	x, y := i.Current()
	if x != 2 && y != 2 {
		t.Errorf("Expected position at (2, 2), got (%d, %d) instead.", x, y)
	}
}

func TestIterator_Move(t *testing.T) {
	i := NewIterator()
	i.MoveOne()
	x, y := i.Current()
	if x != 1 && y != 0 {
		t.Errorf("Expected position at (1, 0), got (%d, %d) instead.", x, y)
	}
}

func TestIterator_DistanceFromCentre(t *testing.T) {
	for _, s := range doubles {
		i := NewIteratorAtPosition(s.p.Coordinates())
		d := i.DistanceFromCentre()
		if d != s.d {
			t.Errorf("Expected distance of %d, got %d instead", s.d, d)
		}
	}
}
