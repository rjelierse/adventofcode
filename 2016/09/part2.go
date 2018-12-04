package day09

import (
	"bytes"
	"fmt"
)

type part2 []byte

func (src part2) DecompressedLength() int {
	return expandedLength(src)
}

func extractMarker(src []byte, pos int) (next, count, times int) {
	var start, end int
	start = pos + 1
	end = start + bytes.IndexByte(src[start:], ')')
	marker := src[start:end]
	if _, err := fmt.Sscanf(string(marker), "%dx%d", &count, &times); err != nil {
		panic(err)
	}
	next = end + 1
	return
}

func expandedLength(src []byte) (length int) {
	for i := 0; i < len(src); {
		switch src[i] {
		case '(':
			start, count, times := extractMarker(src, i)
			end := start + count
			seq := src[start:end]
			tot := bytes.Repeat(seq, times)
			i = end
			length += expandedLength(tot)
		default:
			i++
			length++
		}
	}
	return
}
