package day08

import "fmt"

type Instruction interface {
	Apply(s Screen)
}

type FillInstruction struct {
	width int
	height int
}

type RotateInstruction struct {
	index int
	offset int
}

type RotateColumnInstruction struct {
	RotateInstruction
}

type RotateRowInstruction struct {
	RotateInstruction
}

func NewFillInstruction(line string) Instruction {
	i := &FillInstruction{}
	if _, err := fmt.Sscanf(line, "rect %dx%d", &i.width, &i.height); err != nil {
		panic(err)
	}
	return i
}

func (i FillInstruction) Apply(s Screen) {
	s.Rect(i.width, i.height)
}

func NewRotateColumnInstruction(line string) Instruction {
	i := &RotateColumnInstruction{}
	if _, err := fmt.Sscanf(line, "rotate column x=%d by %d", &i.index, &i.offset); err != nil {
		panic(err)
	}
	return i
}

func (i RotateColumnInstruction) Apply(s Screen) {
	s.RotateColumn(i.index, i.offset)
}

func NewRotateRowInstruction(line string) Instruction {
	i := &RotateRowInstruction{}
	if _, err := fmt.Sscanf(line, "rotate row y=%d by %d", &i.index, &i.offset); err != nil {
		panic(err)
	}
	return i
}

func (i RotateRowInstruction) Apply(s Screen) {
	s.RotateRow(i.index, i.offset)
}
