package day04

import "testing"

func TestRoom_IsValid(t *testing.T) {
	if !NewRoom([]byte("aaaaa-bbb-z-y-x-123[abxyz]")).IsValid() {
		t.Error("Expected room to be valid")
	}
	if !NewRoom([]byte("a-b-c-d-e-f-g-h-987[abcde]")).IsValid() {
		t.Error("Expected room to be valid")
	}
	if !NewRoom([]byte("not-a-real-room-404[oarel]")).IsValid() {
		t.Error("Expected room to be valid")
	}
	if NewRoom([]byte("totally-real-room-200[decoy]")).IsValid() {
		t.Error("Expected room to be invalid")
	}
}
