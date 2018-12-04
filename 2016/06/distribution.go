package day06

type CharacterDistribution map[int]map[byte]int

func NewCharacterDistributionFromMessages(messages [][]byte) CharacterDistribution {
	counts := make(CharacterDistribution)
	for _, message := range messages {
		for col, char := range message {
			if _, exists := counts[col]; !exists {
				counts[col] = make(map[byte]int)
			}
			counts[col][char]++
		}
	}
	return counts
}

func (m CharacterDistribution) MostCommon() []byte {
	message := make([]byte, len(m))
	for col, chars := range m {
		var max int
		for char, count := range chars {
			if count > max {
				max = count
				message[col] = char
			}
		}
	}
	return message
}

func (m CharacterDistribution) LeastCommon() []byte {
	message := make([]byte, len(m))
	for col, chars := range m {
		var min int
		for char, count := range chars {
			if min == 0 || count < min {
				min = count
				message[col] = char
			}
		}
	}
	return message
}

