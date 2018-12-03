package input

func ReadAsStrings(path string) []string {
	data := ReadAsSplitLines(path)
	var strings []string
	for _, item := range data {
		strings = append(strings, string(item))
	}
	return strings
}