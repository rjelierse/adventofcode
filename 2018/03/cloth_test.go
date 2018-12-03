package day03

import "testing"

func TestCloth_CountOverlaps(t *testing.T) {
	cloth := NewCloth()
	cloth.AddClaim(&Claim{1, 1, 3, 4, 4})
	cloth.AddClaim(&Claim{2, 3, 1, 4, 4})
	cloth.AddClaim(&Claim{3, 5, 5, 2, 2})
	count := cloth.CountOverlaps()
	if count != 4 {
		t.Errorf("Unexpected number of overlapping tiles. Expected 4, got %d", count)
	}
}
