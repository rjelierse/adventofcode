package day04

import (
	"bytes"
	"regexp"
	"strconv"
)

var roomExpression = regexp.MustCompile("^([a-z\\-]+)-([0-9]+)\\[([a-z]+)]$")

type Room struct {
	encryptedName []byte
	ID int
	Checksum []byte
}

func NewRoom(input []byte) *Room {
	matches := roomExpression.FindSubmatch(input)
	if matches == nil {
		return nil
	}
	id, _ := strconv.Atoi(string(matches[2]))
	return &Room{
		encryptedName: matches[1],
		ID: id,
		Checksum: matches[3],
	}
}

func (r *Room) Name() string {
	decrypt := make([]byte, len(r.encryptedName))
	for i, char := range r.encryptedName {
		switch char {
		case '-':
			decrypt[i] = ' '
		default:
			decrypt[i] = shift(char, r.ID)
		}
	}
	return string(decrypt)
}

func (r *Room) IsValid() bool {
	return bytes.Equal(r.Checksum, r.calculateChecksum())
}

func (r *Room) calculateChecksum() []byte {
	counts := NewCounterMap(r.encryptedName).Values()
	checksum := make([]byte, 0, len(counts))
	for _, counter := range counts {
		checksum = append(checksum, counter.Character)
	}
	return checksum[:5]
}

func shift(char byte, offset int) byte {
	v := int(char)
	v -= 97
	v += offset
	v %= 26
	v += 97
	return byte(v)
}
