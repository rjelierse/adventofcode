package day03

import "testing"

func TestNewClaim(t *testing.T) {
	c := NewClaim("#1 @ 3,2: 4x5")
	if c.id != 1 {
		t.Errorf("Expected claim to have ID 1, got %d instead", c.id)
	}
	if c.top != 2 {
		t.Errorf("Expected claim to have top-most coordinate 2, got %d instead", c.top)
	}
	if c.left != 3 {
		t.Errorf("Expected claim to have left-most coordinate at 3, got %d instead", c.left)
	}
	if c.width != 4 {
		t.Errorf("Expected claim to have width of 4, got %d instead", c.width)
	}
	if c.height != 5 {
		t.Errorf("Expected claim to have height of 5, got %d instead", c.height)
	}
}

func TestClaim_Occupies(t *testing.T) {
	c := Claim{1, 1, 3, 4, 4}
	o := c.Occupies()

	if len(o) != 16 {
		t.Errorf("Unexpected number of tiles. Expected 16, got %d.", len(o))
	}
}
