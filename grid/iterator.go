package grid

type Iterator struct {
	Position  Position
	Direction Direction
}

func NewIterator() *Iterator {
	return NewIteratorAtPosition(0, 0)
}

func NewIteratorAtPosition(x, y int) *Iterator {
	return NewIteratorAtPositionWithDirection(x, y, DirectionRight)
}

func NewIteratorAtPositionWithDirection(x, y int, dir Direction) *Iterator {
	return &Iterator{
		Position: Position{x, y},
		Direction: dir,
	}
}

func (i *Iterator) Current() (int, int) {
	return i.Position.Coordinates()
}

func (i *Iterator) TurnLeft() {
	i.Direction = i.Direction.TurnLeft()
}

func (i *Iterator) TurnRight() {
	i.Direction = i.Direction.TurnRight()
}

func (i *Iterator) MoveOne() {
	i.Move(1)
}

func (i *Iterator) Move(distance int) {
	i.Position = i.Position.Move(i.Direction, distance)
}

func (i *Iterator) DistanceFromCentre() int {
	return i.DistanceFromPosition(Position{0, 0})
}

func (i *Iterator) DistanceFromPosition(p Position) int {
	ax, ay := i.Current()
	px, py := p.Coordinates()
	var dx, dy int
	if ax > px {
		dx = ax - px
	} else {
		dx = px - ax
	}
	if ay > py {
		dy = ay - py
	} else {
		dy = py - ay
	}
	return dx + dy
}
