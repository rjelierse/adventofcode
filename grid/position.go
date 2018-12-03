package grid

type Position [2]int

func (p Position) Coordinates() (int, int) {
	return p[0], p[1]
}

func (p Position) Move(dir Direction, dist int) Position {
	x, y := p.Coordinates()
	switch dir {
	case DirectionUp:
		return Position{x, y - dist}
	case DirectionDown:
		return Position{x, y + dist}
	case DirectionLeft:
		return Position{x - dist, y}
	case DirectionRight:
		return Position{x + dist, y}
	default:
		panic("unknown direction")
	}
}

func (p Position) MoveOne(dir Direction) Position {
	return p.Move(dir, 1)
}

func (p Position) InBounds(xMin, xMax, yMin, yMax int) bool {
	x, y := p.Coordinates()
	return xMin <= x && x <= xMax && yMin <= y && y <= yMax
}
