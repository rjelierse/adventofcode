package day24

import (
	"context"
	"flag"
	"github.com/google/subcommands"
	"io/ioutil"
	"fmt"
	"os"
	"strings"
	"sort"
)

type command struct{
	path string
}

func (c *command) Name() string {
	return "day24"
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
	input, err := ioutil.ReadFile(c.path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot read input:", err)
		return subcommands.ExitFailure
	}
	links, err := ParseLinks(strings.TrimSpace(string(input)))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Cannot parse input:", err)
		return subcommands.ExitFailure
	}
	bridges := links.RecursiveSearch(0)
	sort.Sort(bridges)
	sort.Reverse(bridges)
	fmt.Println("Strongest bridge:", bridges[0].Strength())
	return subcommands.ExitSuccess
}

func Command() subcommands.Command {
	return &command{}
}
