package input

import (
	"strconv"
)

// ReadAsIntSlice gathers the input data and returns the integer values for each line
func ReadAsIntSlice(path string) []int {
	data := ReadAsSplitLines(path)
	var ints []int

	for _, item := range data {
		value, err := strconv.Atoi(string(item))
		if err != nil {
			panic(err)
		}
		ints = append(ints, value)
	}

	return ints
}

