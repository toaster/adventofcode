package main

import (
	"fmt"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	lines := io.ReadLines()
	directions := lines[0]
	connections := map[string][2]string{}
	for _, line := range lines[2:] {
		srcAndDest := strings.Split(line, " = ")
		connections[srcAndDest[0]] = [2]string{srcAndDest[1][1:4], srcAndDest[1][6:9]}
	}
	cur := "AAA"
	steps := 0
	for {
		for _, d := range directions {
			if d == 'L' {
				cur = connections[cur][0]
			} else {
				cur = connections[cur][1]
			}
			steps++
			if cur == "ZZZ" {
				fmt.Println(steps)
				return
			}
		}
	}
}
