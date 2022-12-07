package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

type move struct {
	count  int
	source int
	dest   int
}

func main() {
	input := io.ReadLines()
	stacks, moves := parseStacksAndMoves(input)

	for _, m := range moves {
		cutIndex := len(stacks[m.source]) - m.count
		stacks[m.dest] = append(stacks[m.dest], stacks[m.source][cutIndex:]...)
		stacks[m.source] = stacks[m.source][:cutIndex]
	}
	for _, s := range stacks {
		fmt.Print(string(s[len(s)-1]))
	}
	fmt.Println()
}

func parseStacksAndMoves(input []string) ([][]rune, []move) {
	separatorIndex := 0
	for i, line := range input {
		if line == "" {
			separatorIndex = i
			break
		}
	}
	stackCount := (len(input[separatorIndex-1]) + 2) / 4
	stacks := make([][]rune, stackCount)
	for i := separatorIndex - 2; i >= 0; i-- {
		for j := 0; j < stackCount; j++ {
			if j*4 > len(input[i]) {
				continue
			}
			v := input[i][j*4+1]
			if v == ' ' {
				continue
			}
			stacks[j] = append(stacks[j], rune(v))
		}
	}
	var moves []move
	for _, line := range input[separatorIndex+1:] {
		values := strings.Split(line, " ")
		m := move{
			count:  io.ParseInt(values[1]),
			source: io.ParseInt(values[3]) - 1,
			dest:   io.ParseInt(values[5]) - 1,
		}
		moves = append(moves, m)
	}
	return stacks, moves
}

func printStacks(stacks [][]rune) {
	height := 0
	for _, s := range stacks {
		if height < len(s) {
			height = len(s)
		}
	}
	for i := height - 1; i >= 0; i-- {
		for _, s := range stacks {
			if len(s) > i {
				fmt.Print(string(s[i]))
			} else {
				fmt.Print(" ")
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println()
}
