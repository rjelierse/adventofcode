package day09

import (
	"bytes"
)

type part1 []byte

func (src part1) DecompressedLength() int {
	return len(src.Decompress())
}

func (src part1) Decompress() (dst []byte) {
	for i := 0; i < len(src); {
		switch src[i] {
		case '(':
			start, count, times := extractMarker(src, i)
			end := start + count
			seq := src[start:end]
			tot := bytes.Repeat(seq, times)
			dst = append(dst, tot...)
			i = end
		default:
			dst = append(dst, src[i])
			i++
		}
	}
	return
}
