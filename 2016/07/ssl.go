package day07

import "bytes"

func isABAsequence(seq []byte) bool {
	if bytes.ContainsAny(seq, "[]") {
		return false
	}
	return seq[0] == seq[2] && seq[0] != seq[1]
}

func hasBABsequence(addr, bab []byte) bool {
	max := len(addr) - 2
	var inHypernet bool
	for i := 0; i < max; i++ {
		switch addr[i] {
		case '[':
			inHypernet = true
		case ']':
			inHypernet = false
		default:
			if !inHypernet {
				continue
			}
			seq := addr[i : i+3]
			if bytes.Equal(bab, seq) {
				return true
			}
		}
	}
	return false
}

func SupportsSSL(addr []byte) bool {
	max := len(addr) - 2
	var inHypernet bool
	for i := 0; i < max; i++ {
		switch addr[i] {
		case '[':
			inHypernet = true
		case ']':
			inHypernet = false
		default:
			if inHypernet {
				continue
			}
			seq := addr[i : i+3]
			if !isABAsequence(seq) {
				continue
			}
			rev := []byte{seq[1], seq[0], seq[1]}
			if hasBABsequence(addr, rev) {
				return true
			}
		}
	}
	return false
}
