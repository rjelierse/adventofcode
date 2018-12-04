package day07

import "testing"

var tlsAddresses = map[string]bool{
	"abba[mnop]qrst":       true,
	"abcd[bddb]xyyx":       false,
	"aaaa[qwer]tyui":       false,
	"ioxxoj[asdfgh]zxcvbn": true,
}

func TestSupportsTLS(t *testing.T) {
	for address, expected := range tlsAddresses {
		actual := SupportsTLS([]byte(address))
		if expected != actual {
			t.Errorf("For address '%s', expected %t, got %t.", address, expected, actual)
		}
	}
}
