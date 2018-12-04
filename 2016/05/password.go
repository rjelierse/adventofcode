package day05

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

type Password string

func NewPassword(prefix []byte, count int) Password {
	password := make([]byte, len(prefix))
	copy(password, prefix)
	suffix := []byte(strconv.Itoa(count))
	password = append(password, suffix...)
	return Password(fmt.Sprintf("%x", md5.Sum(password)))
}

func (p Password) IsNext() bool {
	return p[:5] == "00000"
}

func (p Password) NextCharacter() (c byte) {
	c = byte(p[5])
	return
}

func (p Password) NextCharacterAtPosition() (pos int, c byte) {
	pos = int(p[5]) - 48
	c = byte(p[6])
	return
}
