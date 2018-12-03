package input

import "bytes"

func ReadAsSplitLines(path string) [][]byte {
	data := ReadAsIs(path)
	return bytes.Split(data, []byte{'\n'})
}
