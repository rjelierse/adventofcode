package day09

import "testing"

func TestPart2_DecompressedLength(t *testing.T) {
	doubles := map[string]int {
		"(3x3)XYZ": 9,
		"X(8x2)(3x3)ABCY": 20,
		"(27x12)(20x12)(13x14)(7x10)(1x12)A": 241920,
		"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN": 445,
	}
	for data, expected := range doubles {
		actual := part2(data).DecompressedLength()
		if expected != actual {
			t.Errorf("Expected decompressed length of %d, got %d instead.", expected, actual)
		}
	}
}
