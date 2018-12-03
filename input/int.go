package input

import (
	"log"
	"strconv"
)

// ReadAsInts gathers the input data and returns the integer values for each line
func ReadAsInts(path string) []int {
	data := ReadAsSplitLines(path)
	var ints []int

	for _, item := range data {
		value, err := strconv.Atoi(string(item))
		if err != nil {
			log.Fatal("Cannot parse number:", err)
		}
		ints = append(ints, value)
	}

	return ints
}

