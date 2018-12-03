package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode/2017/01"
	"github.com/rjelierse/adventofcode/2017/02"
	"github.com/rjelierse/adventofcode/2017/03"
	"github.com/rjelierse/adventofcode/2017/04"
	"github.com/rjelierse/adventofcode/2017/05"
	"github.com/rjelierse/adventofcode/2017/06"
	"github.com/rjelierse/adventofcode/2017/07"
	"github.com/rjelierse/adventofcode/2017/08"
	"github.com/rjelierse/adventofcode/2017/09"
	"github.com/rjelierse/adventofcode/2017/10"
	"github.com/rjelierse/adventofcode/2017/11"
	"github.com/rjelierse/adventofcode/2017/12"
	"github.com/rjelierse/adventofcode/2017/13"
	"github.com/rjelierse/adventofcode/2017/14"
	"github.com/rjelierse/adventofcode/2017/15"
	"github.com/rjelierse/adventofcode/2017/16"
	"github.com/rjelierse/adventofcode/2017/17"
	"github.com/rjelierse/adventofcode/2017/18"
	"github.com/rjelierse/adventofcode/2017/19"
	"github.com/rjelierse/adventofcode/2017/20"
	"github.com/rjelierse/adventofcode/2017/21"
	"github.com/rjelierse/adventofcode/2017/22"
	"github.com/rjelierse/adventofcode/2017/23"
	"github.com/rjelierse/adventofcode/2017/24"
	"github.com/rjelierse/adventofcode/2017/25"
	"os"
)

func main() {
	// Register default commands
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.HelpCommand(), "")

	subcommands.Register(day01.Command(), "Puzzles")
	subcommands.Register(day02.Command(), "Puzzles")
	subcommands.Register(day03.Command(), "Puzzles")
	subcommands.Register(day04.Command(), "Puzzles")
	subcommands.Register(day05.Command(), "Puzzles")
	subcommands.Register(day06.Command(), "Puzzles")
	subcommands.Register(day07.Command(), "Puzzles")
	subcommands.Register(day08.Command(), "Puzzles")
	subcommands.Register(day09.Command(), "Puzzles")
	subcommands.Register(day10.Command(), "Puzzles")
	subcommands.Register(day11.Command(), "Puzzles")
	subcommands.Register(day12.Command(), "Puzzles")
	subcommands.Register(day13.Command(), "Puzzles")
	subcommands.Register(day14.Command(), "Puzzles")
	subcommands.Register(day15.Command(), "Puzzles")
	subcommands.Register(day16.Command(), "Puzzles")
	subcommands.Register(day17.Command(), "Puzzles")
	subcommands.Register(day18.Command(), "Puzzles")
	subcommands.Register(day19.Command(), "Puzzles")
	subcommands.Register(day20.Command(), "Puzzles")
	subcommands.Register(day21.Command(), "Puzzles")
	subcommands.Register(day22.Command(), "Puzzles")
	subcommands.Register(day23.Command(), "Puzzles")
	subcommands.Register(day24.Command(), "Puzzles")
	subcommands.Register(day25.Command(), "Puzzles")

	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
