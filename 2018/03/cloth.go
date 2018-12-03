package day03

import (
	"github.com/rjelierse/adventofcode/grid"
)

type Cloth map[grid.Position]int

func NewCloth() Cloth {
	return make(Cloth)
}

func (c Cloth) AddClaim(claim *Claim) {
	for _, coordinate := range claim.Occupies() {
		if count, exists := c[coordinate]; exists {
			c[coordinate] = count + 1
		} else {
			c[coordinate] = 1
		}
	}
}

func (c Cloth) CountOverlaps() (overlaps int) {
	for _, claims := range c {
		if claims >= 2 {
			overlaps++
		}
	}
	return
}

func (c Cloth) IsUnique(claim *Claim) bool {
	for _, coordinate := range claim.Occupies() {
		if c[coordinate] >= 2 {
			return false
		}
	}
	return true
}
