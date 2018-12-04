package day04

import (
	"fmt"
	"regexp"
	"sort"
	"time"
)

// LogEntry registers a specific action at a specific time
type LogEntry struct {
	When time.Time
	action string
}

// LogEntrySlice contains all LogEntries
type LogEntrySlice []*LogEntry

// LogAction enumerates the possible actions for a LogEntry
type LogAction int

const (
	ShiftStarted LogAction = iota
	FallsAsleep
	WakesUp
)

// NewLogEntrySlice parses the log file lines into entries and returns them in chronological order
func NewLogEntrySlice(lines []string) LogEntrySlice {
	entries := make(LogEntrySlice, 0, len(lines))
	for _, line := range lines {
		entries = append(entries, NewLogEntry(line))
	}
	sort.Sort(entries)
	return entries
}

func (s LogEntrySlice) Len() int {
	return len(s)
}

func (s LogEntrySlice) Less(i, j int) bool {
	return s[i].When.Before(s[j].When)
}

func (s LogEntrySlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

var logRegex = regexp.MustCompile("^\\[([0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2})] (.+)$")

// NewLogEntry parses a single log file line into a LogEntry
func NewLogEntry(line string) *LogEntry {
	matches := logRegex.FindStringSubmatch(line)
	if matches == nil {
		panic("Did not find matches for line")
	}
	when, err := time.Parse("2006-01-02 15:04", matches[1])
	if err != nil {
		panic(err)
	}
	return &LogEntry{
		When: when,
		action: matches[2],
	}
}

// Action returns the LogAction that this entry relates to
func (l *LogEntry) Action() LogAction {
	switch l.action {
	case "falls asleep":
		return FallsAsleep
	case "wakes up":
		return WakesUp
	default:
		return ShiftStarted
	}
}

// Guard determines the guard that began their shift
func (l *LogEntry) Guard() (id int) {
	_, err := fmt.Sscanf(l.action,"Guard #%d begins shift", &id)
	if err != nil {
		return -1
	}
	return
}
