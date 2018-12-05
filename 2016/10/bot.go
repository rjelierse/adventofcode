package day10

import (
	"regexp"
	"strconv"
)

var recvRegex = regexp.MustCompile("^value ([0-9]+) goes to bot ([0-9]+)$")
var giveRegex = regexp.MustCompile("^bot ([0-9]+) gives low to (output|bot) ([0-9]+) and high to (output|bot) ([0-9]+)$")

type BotRepository map[int]*Bot

func (r BotRepository) Get(id int) *Bot {
	_, exists := r[id]
	if !exists {
		r[id] = NewBot(id)
	}
	return r[id]
}

func (r BotRepository) AllReady(bot chan *Bot) {
	for _, b := range r {
		if b.CanProceed() {
			bot <- b
		}
	}
}

type Bot struct {
	ID int
	instruction string
	chips []int
}

func NewBot(id int) *Bot {
	return &Bot{
		ID: id,
		instruction: "",
		chips: make([]int, 0, 2),
	}
}

func (b *Bot) CanProceed() bool {
	return len(b.chips) == 2
}

func (b *Bot) AddChip(chip int) {
	b.chips = append(b.chips, chip)
}

func (b *Bot) Chips() (low, high int) {
	if b.chips[0] < b.chips[1] {
		return b.chips[0], b.chips[1]
	} else {
		return b.chips[1], b.chips[0]
	}
}

func (b *Bot) AddInstruction(i string) {
	b.instruction = i
}

func (b *Bot) Instructions() (lowType string, lowTarget int, highType string, highTarget int) {
	matches := giveRegex.FindStringSubmatch(b.instruction)
	if matches == nil {
		panic("Could not match instructions")
	}
	lowType = matches[2]
	lowTarget, _ = strconv.Atoi(matches[3])
	highType = matches[4]
	highTarget, _ = strconv.Atoi(matches[5])
	return
}
