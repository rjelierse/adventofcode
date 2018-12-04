package day04

import (
	"time"
)

// Guard contains the sleep information for a specific guard
type Guard struct {
	ID     int
	Sleeps []int
	Slept  int
}

// NewGuard creates a new Guard instance
func NewGuard(id int) *Guard {
	return &Guard{
		ID: id,
		Sleeps: make([]int, 60),
		Slept: 0,
	}
}

// RegisterSleep adds a sleep period for the guard
func (g *Guard) RegisterSleep(start, end time.Time) {
	from, to := start.Minute(), end.Minute()
	for i := from; i < to; i++ {
		g.Sleeps[i]++
	}
	g.Slept += to - from
}

// MostSleptMinute determines at which minute the guard was most frequently asleep
func (g *Guard) MostSleptMinute() (minute, max int) {
	for m, count := range g.Sleeps {
		if count > max {
			minute = m
			max = count
		}
	}
	return
}

// GuardMap contains all Guard instances
type GuardMap map[int]*Guard

// NewGuardMap parses the entries into a per-guard overview of when guards were asleep and how much they slept
func NewGuardMap(entries LogEntrySlice) GuardMap {
	var activeGuard *Guard
	var startsSleeping time.Time
	m := make(GuardMap)
	for _, entry := range entries {
		switch entry.Action() {
		case ShiftStarted:
			activeGuard = m.get(entry.Guard())
		case FallsAsleep:
			startsSleeping = entry.When
		case WakesUp:
			activeGuard.RegisterSleep(startsSleeping, entry.When)
		}
	}
	return m
}

func (m GuardMap) get(id int) *Guard {
	_, exists := m[id]
	if !exists {
		m[id] = NewGuard(id)
	}
	return m[id]
}

// MostAsleep returns the guard that catches the most sleep during its shift
func (m GuardMap) MostAsleep() (g *Guard) {
	for _, v := range m {
		if g == nil {
			g = v
		}
		if v.Slept > g.Slept {
			g = v
		}
	}
	return
}

// MostSleptAtSameTime returns the guard that catches the most sleep at the same minute
func (m GuardMap) MostSleptAtSameTime() (g *Guard) {
	var count int
	for _, v := range m {
		_, c := v.MostSleptMinute()
		if c > count {
			count = c
			g = v
		}
	}
	return
}
