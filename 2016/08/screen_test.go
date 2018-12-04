package day08

import "testing"

func TestScreen_Count(t *testing.T) {
	s := NewScreen(7, 3)
	s.Rect(3, 2)
	s.RotateColumn(1, 1)
	s.RotateRow(0, 4)
	s.RotateColumn(1, 1)
	c := s.Count()
	if c != 6 {
		t.Errorf("Expected 6 lit pixels, got %d instead.", c)
	}
}
