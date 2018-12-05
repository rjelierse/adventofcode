package day05

type Polymer []byte

const capOffset = 'a' - 'A'

func NewPolymer(data []byte) Polymer {
	p := make(Polymer, len(data))
	copy(p, data)
	return p
}

func (p Polymer) Len() int {
	return len(p)
}

func (p Polymer) Fold() Polymer {
	c := NewPolymer(p)
	for i := 0; i < len(c) - 1; {
		switch  {
		case canReact(c[i], c[i+1]):
			c = append(c[:i], c[i+2:]...)
			if i > 0 {
				i--
			}
		default:
			i++
		}
	}
	return c
}

func (p Polymer) Filter(char byte) Polymer {
	c := NewPolymer(p)
	for i := 0; i < len(c); {
		switch {
		case shouldRemove(c[i], char):
			c = append(c[:i], c[i+1:]...)
		default:
			i++
		}
	}
	return c
}

func canReact(a, b byte) bool {
	return a - b == capOffset || b - a == capOffset
}

func shouldRemove(current, matches byte) bool {
	return current == matches || current == matches + capOffset
}
