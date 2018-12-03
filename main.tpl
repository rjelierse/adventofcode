package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode/YEAR/01"
	"github.com/rjelierse/adventofcode/YEAR/02"
	"github.com/rjelierse/adventofcode/YEAR/03"
	"github.com/rjelierse/adventofcode/YEAR/04"
	"github.com/rjelierse/adventofcode/YEAR/05"
	"github.com/rjelierse/adventofcode/YEAR/06"
	"github.com/rjelierse/adventofcode/YEAR/07"
	"github.com/rjelierse/adventofcode/YEAR/08"
	"github.com/rjelierse/adventofcode/YEAR/09"
	"github.com/rjelierse/adventofcode/YEAR/10"
	"github.com/rjelierse/adventofcode/YEAR/11"
	"github.com/rjelierse/adventofcode/YEAR/12"
	"github.com/rjelierse/adventofcode/YEAR/13"
	"github.com/rjelierse/adventofcode/YEAR/14"
	"github.com/rjelierse/adventofcode/YEAR/15"
	"github.com/rjelierse/adventofcode/YEAR/16"
	"github.com/rjelierse/adventofcode/YEAR/17"
	"github.com/rjelierse/adventofcode/YEAR/18"
	"github.com/rjelierse/adventofcode/YEAR/19"
	"github.com/rjelierse/adventofcode/YEAR/20"
	"github.com/rjelierse/adventofcode/YEAR/21"
	"github.com/rjelierse/adventofcode/YEAR/22"
	"github.com/rjelierse/adventofcode/YEAR/23"
	"github.com/rjelierse/adventofcode/YEAR/24"
	"github.com/rjelierse/adventofcode/YEAR/25"
)

func main() {
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
	subcommands.Execute(ctx)
}
