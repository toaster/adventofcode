package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	notes := parseNotes(io.ReadLines())
	count := 0
	for _, n := range notes {
		for _, out := range n.output {
			switch len(out) {
			case 2, 3, 4, 7: // 1, 7, 4, 8
				count++
			}
		}
	}
	fmt.Println(count)
}

func parseNotes(lines []string) (notes []*note) {
	for _, line := range lines {
		raw := strings.Split(line, " | ")
		notes = append(notes, &note{
			signals: strings.Split(raw[0], " "),
			output:  strings.Split(raw[1], " "),
		})
	}
	return
}

type note struct {
	signals []string
	output  []string
}
