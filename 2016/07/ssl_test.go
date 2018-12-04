package day07

import (
	"testing"
)

var sslAddresses = map[string]bool{
	"aba[bab]xyz":   true,
	"xyx[xyx]xyx":   false,
	"aaa[kek]eke":   true,
	"zazbz[bzb]cdb": true,
}

func TestSupportsSSL(t *testing.T) {
	for address, expected := range sslAddresses {
		actual := SupportsSSL([]byte(address))
		if expected != actual {
			t.Errorf("For address '%s', expected %t, got %t.", address, expected, actual)
		}
	}
}
