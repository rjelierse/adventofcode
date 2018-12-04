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

// ReadAsInt returns the input data as an integer value
func ReadAsInt(path string) int {
	data := ReadAsIs(path)
	value, err := strconv.Atoi(string(data))
	if err != nil {
		panic(err)
	}
	return value
}
