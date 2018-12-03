package day23

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
	"io/ioutil"
	"os"
	"strings"
)

type command struct {
	path string
}

func (c *command) Name() string {
	return "day23"
}

func (c *command) Synopsis() string {
	return ""
}

func (c *command) Usage() string {
	return ""
}

func (c *command) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.path, "input", "", "")
}

func (c *command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	data, err := ioutil.ReadFile(c.path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot read input:", err)
		return subcommands.ExitFailure
	}
	input := strings.Split(strings.TrimSpace(string(data)), "\n")
	fmt.Println("Number of 'mul' instructions:", program(input))
	fmt.Println("Value of 'h':", run(79))
	return subcommands.ExitSuccess
}

func Command() subcommands.Command {
	return &command{}
}

// run is the input assembly written as Go code (with a few optimizations)
func run(b int) int {
	var c, g, h int
	b = (b * 100) + 100000
	c = b + 17000
	for {
		for d := 2; d * d <= b; d++ {
			if b % d == 0 {
				h++
				break
			}
		}
		g = b - c
		b += 17
		if g == 0 {
			break
		}
	}
	return h
}
