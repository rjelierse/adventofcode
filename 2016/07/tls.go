package day07

import "bytes"

func isABBAsequence(seq []byte) bool {
	if bytes.ContainsAny(seq, "[]") {
		return false
	}
	if seq[0] == seq[1] || seq[2] == seq[3] {
		return false
	}
	return seq[0] == seq[3] && seq[1] == seq[2]
}

func SupportsTLS(address []byte) (supported bool) {
	max := len(address) - 3
	var inHypernet bool
	for i := 0; i < max; i++ {
		switch address[i] {
		case '[':
			inHypernet = true
		case ']':
			inHypernet = false
		default:
			isABBA := isABBAsequence(address[i : i+4])
			// When ABBA sequence is found within hypernet, does not support TLS
			if isABBA && inHypernet {
				return false
			}
			if isABBA {
				supported = true
			}
		}
	}
	return
}
