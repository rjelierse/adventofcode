package input

import (
	"bytes"
	"io/ioutil"
)

func ReadAsIs(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return bytes.TrimSpace(data)
}
