package input

func ReadAsStringSlice(path string) []string {
	data := ReadAsSplitLines(path)
	var strings []string
	for _, item := range data {
		strings = append(strings, string(item))
	}
	return strings
}

func ReadAsString(path string) string {
	return string(ReadAsIs(path))
}
