package day08

const (
	OperatorIncrease Operator = "inc"
	OperatorDecrease          = "dec"
)

type Operator string

type Instruction struct {
	Register string
	Operator Operator
	Value    int
}
