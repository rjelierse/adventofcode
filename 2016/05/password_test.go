package day05

import "testing"

func TestPassword_NextCharacter(t *testing.T) {
	p := NewPassword([]byte("abc"), 3231929)
	if !p.IsNext() {
		t.Errorf("Expected valid hash")
	}
	c := p.NextCharacter()
	if c != '1' {
		t.Errorf("Expected 6th character in hash to be '1', got '%c' instead.", c)
	}
}

func TestPassword_NextCharacterAtPosition(t *testing.T) {
	p := NewPassword([]byte("abc"), 3231929)
	if !p.IsNext() {
		t.Errorf("Expected valid hash")
	}
	pos, c := p.NextCharacterAtPosition()
	if pos != 1 {
		t.Errorf("Expected 6th character in hash to be '1', got '%d' instead.", pos)
	}
	if c != '5' {
		t.Errorf("Expected 7th character in hash to be '5', got '%c' instead.", c)
	}
}
