package day24

import (
	"testing"
	"sort"
)

const linksInput = "0/2\n2/2\n2/3\n3/4\n3/5\n0/1\n10/1\n9/10"

func TestParseLinks(t *testing.T) {
	pool, err := ParseLinks(linksInput)
	if err != nil {
		t.Fatal("Failed to parse input:", err)
	}
	if pool.Len() != 8 {
		t.Error("Expected pool of 8 elements, got", pool.Len())
	}
}

func TestPool_RecursiveSearch(t *testing.T) {
	pool, err := ParseLinks(linksInput)
	if err != nil {
		t.Fatal("Failed to parse input:", err)
	}
	bridges := pool.RecursiveSearch(0)
	sort.Sort(bridges)
	sort.Reverse(bridges)
	if bridges[0].Strength() != 31 {
		t.Error("Expected strongest bridge at 31, got", bridges[0].Strength())
	}
}
