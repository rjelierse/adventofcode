package day01

func part1(changes []int) int {
	frequency := 0
	for _, change := range changes {
		frequency = frequency + change
	}
	return frequency
}

