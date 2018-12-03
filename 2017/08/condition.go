package day08

const (
	CompareEquals       Comparator = "=="
	CompareNotEquals               = "!="
	CompareGreater                 = ">"
	CompareGreaterEqual            = ">="
	CompareLower                   = "<"
	CompareLowerEqual              = "<="
)

type Comparator string

type Condition struct {
	Register   string
	Comparator Comparator
	Value      int
}
