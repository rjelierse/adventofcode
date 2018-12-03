package day02

func part2(boxes []string) string {
	for _, box := range boxes {
		for _, otherBox := range boxes {
			var differentCharacterCount int
			var commonCharacters []byte
			for index, char := range []byte(box) {
				if char != otherBox[index] {
					differentCharacterCount = differentCharacterCount + 1
				} else {
					commonCharacters = append(commonCharacters, char)
				}
			}
			if differentCharacterCount == 1 {
				return string(commonCharacters)
			}
		}
	}
	return ""
}
