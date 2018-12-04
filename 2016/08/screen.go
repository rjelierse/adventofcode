package day08

import "fmt"

type Screen struct {
	pixels [][]bool
	width int
	height int
}

func NewScreen(width, height int) Screen {
	screen := make([][]bool, height)
	for i := 0; i < height; i++ {
		screen[i] = make([]bool, width)
	}
	return Screen{
		pixels: screen,
		width: width,
		height: height,
	}
}

func (s Screen) Rect(width, height int) {
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			s.pixels[y][x] = true
		}
	}
}

func (s Screen) RotateColumn(index, offset int) {
	values := make([]bool, s.height)
	for src := 0; src < s.height; src++ {
		dst := (src + offset) % s.height
		values[dst] = s.pixels[src][index]
	}
	for dst, value := range values {
		s.pixels[dst][index] = value
	}
}

func (s Screen) RotateRow(index, offset int) {
	values := make([]bool, s.width)
	for src := 0; src < s.width; src++ {
		dst := (src + offset) % s.width
		values[dst] = s.pixels[index][src]
	}
	for dst, value := range values {
		s.pixels[index][dst] = value
	}
}

func (s Screen) Count() (lit int) {
	for x := 0; x < s.width; x++ {
		for y := 0; y < s.height; y++ {
			if s.pixels[y][x] {
				lit++
			}
		}
	}
	return
}

func (s Screen) Print() {
	for y := 0; y < s.height; y++ {
		for x := 0; x < s.width; x++ {
			if s.pixels[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}