package day24

import (
	"container/heap"
	"strings"
)

type Pool []*Link

func (p Pool) RecursiveSearch(parity int) Bridges {
	bridges := make(Bridges, 0)
	for n, l := range p {
		if l.Contains(parity) {
			bridge := Bridge{l}
			bridges = append(bridges, bridge)
			pool := append(p[0:n], p[n+1:]...)
			for _, b := range pool.RecursiveSearch(l.Other(parity)) {
				bridges = append(bridges, append(bridge, b...))
			}
		}
	}
	return bridges
}

func (p Pool) Len() int {
	return len(p)
}

func (p Pool) Less(i, j int) bool {
	switch {
	case p[i].Contains(0) && p[j].Contains(0):
		return p[i].Sum() < p[j].Sum()
	case p[i].Contains(0):
		return true
	case p[j].Contains(0):
		return false
	default:
		return p[i].Sum() < p[j].Sum()
	}
}

func (p Pool) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *Pool) Push(v interface{}) {
	link := v.(*Link)
	*p = append(*p, link)
}

func (p *Pool) Pop() interface{} {
	old := *p
	n := len(old)
	e := old[n- 1]
	*p = old[0 : n- 1]
	return e
}

func ParseLinks(input string) (*Pool, error) {
	p := &Pool{}
	heap.Init(p)
	for _, line := range strings.Split(input, "\n") {
		l, err := NewLink(line)
		if err != nil {
			return nil, err
		}
		heap.Push(p, l)
	}
	return p, nil
}
