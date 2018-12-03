package day02

func characterCounts(box string) map[int32]int {
	counts := make(map[int32]int)
	for _, char := range box {
		counts[char] = counts[char] + 1
	}
	return counts
}

func part1(boxes []string) int {
	var countDoubles, countTriples int
	for _, box := range boxes {
		counts := characterCounts(box)
		var hasDouble, hasTriple bool
		for _, count := range counts {
			if count == 3 {
				hasTriple = true
			}
			if count == 2 {
				hasDouble = true
			}
		}
		if hasDouble {
			countDoubles = countDoubles + 1
		}
		if hasTriple {
			countTriples = countTriples + 1
		}
	}
	return countDoubles * countTriples
}
