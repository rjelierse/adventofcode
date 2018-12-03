package day23

import (
	"strconv"
	"strings"
)

func program(instructions []string) (muls int) {
	registers := map[string]int{}

	get := func(s string) int {
		if strings.IndexAny(s, "0123456789") != -1 {
			v, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			return v
		}
		return registers[s]
	}

	var pos int
	for pos >= 0 && pos < len(instructions) {
		fields := strings.Fields(instructions[pos])
		switch fields[0] {
		case "jgz":
			if get(fields[1]) > 0 {
				pos += get(fields[2])
				continue
			}
		case "jnz":
			if get(fields[1]) != 0 {
				pos += get(fields[2])
				continue
			}
		case "set":
			registers[fields[1]] = get(fields[2])
		case "add":
			registers[fields[1]] += get(fields[2])
		case "sub":
			registers[fields[1]] -= get(fields[2])
		case "mul":
			registers[fields[1]] *= get(fields[2])
			muls++
		case "mod":
			registers[fields[1]] %= get(fields[2])
		default:
			panic("unknown command")
		}
		pos++
	}

	return muls
}
