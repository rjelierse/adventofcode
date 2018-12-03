package adventofcode

import (
	"context"
	"flag"
	"github.com/google/subcommands"
)

type Puzzle struct {
	name    string
	path    string
	handler func(string) error
}

func NewPuzzle(n string, f func(string) error) *Puzzle {
	return &Puzzle{
		name:    n,
		handler: f,
	}
}

func (p *Puzzle) Name() string {
	return p.name
}

func (*Puzzle) Synopsis() string {
	return ""
}

func (*Puzzle) Usage() string {
	return ""
}

func (p *Puzzle) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.path, "input", "", "The puzzle input")
}

func (p *Puzzle) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if p.path == "" {
		return subcommands.ExitUsageError
	}
	err := p.handler(p.path)
	if err != nil {
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}
