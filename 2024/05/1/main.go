package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	input := io.ReadAll()
	sections := strings.Split(strings.TrimSpace(input), "\n\n")
	rules := map[int][]int{}
	for _, line := range strings.Split(sections[0], "\n") {
		rule := io.ParseInts(line, "|")
		rules[rule[0]] = append(rules[rule[0]], rule[1])
	}
	var updates [][]int
	for _, line := range strings.Split(sections[1], "\n") {
		updates = append(updates, io.ParseInts(line, ","))
	}

	sum := 0
	for _, update := range updates {
		seen := map[int]bool{}
		unordered := false
		for _, page := range update {
			rule := rules[page]
			for _, p := range rule {
				if seen[p] {
					unordered = true
					break
				}
			}
			if unordered {
				break
			}
			seen[page] = true
		}
		if unordered {
			continue
		}
		sum += update[(len(update)-1)/2]
	}
	fmt.Println(sum)
}
