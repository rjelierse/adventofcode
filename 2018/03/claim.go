package day03

import (
	"fmt"
	"github.com/rjelierse/adventofcode/grid"
	"regexp"
	"strconv"
)

type Claim struct {
	id int
	top int
	left int
	width int
	height int
}

var claimExpression = regexp.MustCompile("#([0-9]+) @ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)")

func NewClaim(definition string) *Claim {
	matches := claimExpression.FindStringSubmatch(definition)
	if matches == nil {
		fmt.Println("Invalid input", definition)
		return nil
	}
	id, _ := strconv.Atoi(matches[1])
	top, _ := strconv.Atoi(matches[3])
	left, _ := strconv.Atoi(matches[2])
	width, _ := strconv.Atoi(matches[4])
	height, _ := strconv.Atoi(matches[5])
	return &Claim{
		id: id,
		top: top,
		left: left,
		width: width,
		height: height,
	}
}

func (c *Claim) Occupies() []grid.Position {
	coordinates := make([]grid.Position, c.width * c.height)
	for x := c.left; x < c.left + c.width; x++ {
		for y := c.top; y < c.top + c.height; y++ {
			coordinates = append(coordinates, grid.Position{x, y})
		}
	}
	return coordinates
}
