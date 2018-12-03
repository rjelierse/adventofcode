package main

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"github.com/rjelierse/adventofcode/2016/01"
	"github.com/rjelierse/adventofcode/2016/02"
	"github.com/rjelierse/adventofcode/2016/03"
	"github.com/rjelierse/adventofcode/2016/04"
	"github.com/rjelierse/adventofcode/2016/05"
	"github.com/rjelierse/adventofcode/2016/06"
	"github.com/rjelierse/adventofcode/2016/07"
	"github.com/rjelierse/adventofcode/2016/08"
	"github.com/rjelierse/adventofcode/2016/09"
	"github.com/rjelierse/adventofcode/2016/10"
	"github.com/rjelierse/adventofcode/2016/11"
	"github.com/rjelierse/adventofcode/2016/12"
	"github.com/rjelierse/adventofcode/2016/13"
	"github.com/rjelierse/adventofcode/2016/14"
	"github.com/rjelierse/adventofcode/2016/15"
	"github.com/rjelierse/adventofcode/2016/16"
	"github.com/rjelierse/adventofcode/2016/17"
	"github.com/rjelierse/adventofcode/2016/18"
	"github.com/rjelierse/adventofcode/2016/19"
	"github.com/rjelierse/adventofcode/2016/20"
	"github.com/rjelierse/adventofcode/2016/21"
	"github.com/rjelierse/adventofcode/2016/22"
	"github.com/rjelierse/adventofcode/2016/23"
	"github.com/rjelierse/adventofcode/2016/24"
	"github.com/rjelierse/adventofcode/2016/25"
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
