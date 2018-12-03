package day01

func part2(changes []int) int {
	frequency := 0
	history := make(map[int]bool)
	for {
		for _, change := range changes {
			frequency = frequency + change
			if history[frequency] {
				return frequency
			}
			history[frequency] = true
		}
	}
}

