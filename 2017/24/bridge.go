package day24

type Bridge []*Link
type Bridges []Bridge

func (b Bridge) Strength() int {
	var strength int
	for _, l := range b {
		strength += l.Sum()
	}
	return strength
}

func (b Bridges) Len() int {
	return len(b)
}

func (b Bridges) Less(i, j int) bool {
	return b[i].Strength() < b[j].Strength()
}

func (b Bridges) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
