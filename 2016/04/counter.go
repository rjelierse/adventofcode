package day04

import "sort"

type Counter struct {
	Character byte
	Count     int
}

func NewCounter(char byte) *Counter {
	return &Counter{
		Character: char,
		Count:     1,
	}
}

func (c *Counter) Inc() {
	c.Count++
}

type CounterMap map[byte]*Counter

func NewCounterMap(values []byte) CounterMap {
	counts := make(CounterMap)
	for _, b := range values {
		if b == '-' {
			continue
		}
		if _, exists := counts[b]; !exists {
			counts[b] = NewCounter(b)
		} else {
			counts[b].Inc()
		}
	}
	return counts
}

func (cm CounterMap) Values() []*Counter {
	var values []*Counter
	for _, counter := range cm {
		values = append(values, counter)
	}
	sort.Slice(values, func (i, j int) bool {
		a, b := values[i].Count, values[j].Count
		// if count is equal, sort alphabetically
		if a == b {
			return values[i].Character < values[j].Character
		}
		// higher count should go first
		return values[i].Count > values[j].Count
	})
	return values
}
