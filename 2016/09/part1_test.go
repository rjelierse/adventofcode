package day09

import (
	"bytes"
	"testing"
)

func TestPart1_Decompress(t *testing.T) {
	doubles := map[string][]byte{
		"ADVENT": []byte("ADVENT"),
		"A(1x5)BC": []byte("ABBBBBC"),
		"(3x3)XYZ": []byte("XYZXYZXYZ"),
		"A(2x2)BCD(2x2)EFG": []byte("ABCBCDEFEFG"),
		"(6x1)(1x3)A": []byte("(1x3)A"),
		"X(8x2)(3x3)ABCY": []byte("X(3x3)ABC(3x3)ABCY"),
	}
	for data, expected := range doubles {
		actual := part1(data).Decompress()
		if !bytes.Equal(expected, actual) {
			t.Errorf("Expected %s to decompress to %s, got %s instead.", data, expected, actual)
		}
	}
}

func TestPart1_DecompressedLength(t *testing.T) {
	doubles := map[string]int{
		"ADVENT": 6,
		"A(1x5)BC": 7,
		"(3x3)XYZ": 9,
		"A(2x2)BCD(2x2)EFG": 11,
		"(6x1)(1x3)A": 6,
		"X(8x2)(3x3)ABCY": 18,
	}
	for data, expected := range doubles {
		actual := part1(data).DecompressedLength()
		if expected != actual {
			t.Errorf("Expected decompressed length of %d, got %d instead.", expected, actual)
		}
	}
}