package day24

import (
	"fmt"
)

const linkFormat = "%d/%d"

type Link [2]int

func (l Link) Contains(search int) bool {
	return l[0] == search || l[1] == search
}

func (l Link) Other(parity int) int {
	switch parity{
	case l[0]:
		return l[1]
	case l[1]:
		return l[0]
	default:
		panic("value not found")
	}
}

func (l Link) Sum() int {
	return l[0] + l[1]
}

func (l Link) String() string {
	return fmt.Sprintf(linkFormat, l[0], l[1])
}

func NewLink(input string) (*Link, error) {
	link := Link{}
	_, err := fmt.Sscanf(input, linkFormat, &link[0], &link[1])
	if err != nil {
		return nil, err
	}
	return &link, nil
}
