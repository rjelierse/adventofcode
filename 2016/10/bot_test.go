package day10

import "testing"

var commands = []string{
	"value 5 goes to bot 2",
	"bot 2 gives low to bot 1 and high to bot 0",
	"value 3 goes to bot 1",
	"bot 1 gives low to output 1 and high to bot 0",
	"bot 0 gives low to output 2 and high to output 0",
	"value 2 goes to bot 2",
}

func TestBot_Chips(t *testing.T) {
	bot := &Bot{
		chips: []int{2, 5},
	}
	_, high := bot.Chips()
	if high != 5 {
		t.Errorf("Expected high chip value 5, got %d instead", high)
	}
}

func TestPart1(t *testing.T) {
	id, _ := solve(commands, 2, 5)
	if id != 2 {
		t.Errorf("Expected bot 2 to compare the 2 and 5 chips, got %d instead", id)
	}
}
