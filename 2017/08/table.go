package day08

type Table map[string]int

func NewTable() Table {
	return make(Table)
}

func (t Table) LargestValue() int {
	value := 0

	for _, v := range t {
		if v > value {
			value = v
		}
	}

	return value
}

func (t Table) Matches(c *Condition) bool {
	p := c.Register
	v, exists := t[p]
	if !exists {
		v = 0
	}
	switch c.Comparator {
	case CompareEquals:
		return v == c.Value
	case CompareNotEquals:
		return v != c.Value
	case CompareLower:
		return v < c.Value
	case CompareLowerEqual:
		return v <= c.Value
	case CompareGreaterEqual:
		return v >= c.Value
	case CompareGreater:
		return v > c.Value
	}
	return false
}

func (t Table) Apply(i *Instruction) {
	p := i.Register
	v, exists := t[p]
	if !exists {
		v = 0
	}
	switch i.Operator {
	case OperatorIncrease:
		t[p] = v + i.Value
	case OperatorDecrease:
		t[p] = v - i.Value
	}
}
