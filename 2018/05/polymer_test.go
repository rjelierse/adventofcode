package day05

import "testing"

func TestPolymer_Fold(t *testing.T) {
	result := NewPolymer([]byte("dabAcCaCBAcCcaDA")).Fold()
	if string(result) != "dabCBAcaDA" {
		t.Errorf("Expected result 'dabCBAcaDA', got %s instead", result)
	}
	if result.Len() != 10 {
		t.Errorf("Expected result of length 10, got %d instead", result.Len())
	}
}

func TestPolymer_Filter(t *testing.T) {
	result := NewPolymer([]byte("dabAcCaCBAcCcaDA")).Filter('A')
	if string(result) != "dbcCCBcCcD" {
		t.Errorf("Expected result 'dbcCCBcCcD', got %s instead", result)
	}
}