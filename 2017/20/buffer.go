package day20

import (
	"fmt"
	"sort"
)

type Buffer struct {
	particles  []*Particle
}

func BufferFromInput(input []string) (*Buffer, error) {
	b := new(Buffer)
	b.particles = make([]*Particle, len(input))
	for i, line := range input {
		p, err := ParticleFromInput(i, line)
		if err != nil {
			return nil, err
		}
		b.particles[i] = p
	}
	return b, nil
}

// Closest orders the particles slice by distance at time t and returns the particle closest to <0,0,0>
func (buf *Buffer) Closest(t int) *Particle {
	sort.Slice(buf.particles, func(i, j int) bool {
		return buf.particles[i].CalcPosition(t).Sum() < buf.particles[j].CalcPosition(t).Sum()
	})

	return buf.particles[0]
}

func (buf *Buffer) Run(t int) bool {
	for offset, a := range buf.particles {
		if a.Collided {
			continue
		}
		for _, b := range buf.particles[offset+1:] {
			if b.Collided {
				continue
			}
			if a.CalcPosition(t) == b.CalcPosition(t) {
				fmt.Printf("[%05d] Collision between %d and %d\n", t, a.Id, b.Id)
				a.Collided = true
				b.Collided = true
			}
		}
	}
	return true
}

func (buf *Buffer) Count(collided bool) (result int) {
	for _, p := range buf.particles {
		if p.Collided == collided {
			result++
		}
	}
	return
}
