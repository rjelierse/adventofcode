package day10

import (
	"fmt"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode"
	"github.com/rjelierse/adventofcode/input"
	"strconv"
)

func Command() subcommands.Command {
	return adventofcode.NewPuzzle("day10", func(path string) error {
	    lines := input.ReadAsStringSlice(path)
	    id, mul := solve(lines, 17, 61)
	    fmt.Println("Part 1:", id)
	    fmt.Println("Part 2:", mul)
		return nil
	})
}

func solve(lines []string, a, b int) (id, mul int) {
	outputs := make(map[int]int)
	repository := parseInstructions(lines)
	bots := make(chan *Bot, len(repository))
	go repository.AllReady(bots)
	for bot := range bots {
		lowType, lowID, highType, highID := bot.Instructions()
		low, high := bot.Chips()
		if low == a && high == b {
			id = bot.ID
		}
		bot = nil
		switch lowType {
		case "output":
			outputs[lowID] = low
		case "bot":
			b := repository.Get(lowID)
			b.AddChip(low)
			if b.CanProceed() {
				bots <- b
			}
		}
		switch highType {
		case "output":
			outputs[highID] = high
		case "bot":
			b := repository.Get(highID)
			b.AddChip(high)
			if b.CanProceed() {
				bots <- b
			}
		}
		if lowType == "output" && highType == "output" {
			break
		}
	}
	return id, outputs[0] * outputs[1] * outputs[2]
}

func parseInstructions(lines []string) BotRepository {
	bots := make(BotRepository)
	for _, line := range lines {
		if matches := recvRegex.FindStringSubmatch(line); matches != nil {
			id, _ := strconv.Atoi(matches[2])
			value, _ := strconv.Atoi(matches[1])
			bot := bots.Get(id)
			bot.AddChip(value)
			continue
		}
		if matches := giveRegex.FindStringSubmatch(line); matches != nil {
			id, _ := strconv.Atoi(matches[1])
			bots.Get(id).AddInstruction(line)
			continue
		}
	}
	return bots
}
