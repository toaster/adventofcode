package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	x := 0
	y := 0
	aim := 0
	for _, line := range io.ReadLines() {
		tokens := strings.Split(line, " ")
		if len(tokens) != 2 {
			io.ReportError("failed to parse input", fmt.Errorf("wrong token count: %d", len(tokens)))
		}

		command := tokens[0]
		amount, err := strconv.Atoi(tokens[1])
		io.ReportError("failed to parse input", err)

		switch command {
		case "forward":
			x += amount
			y += aim * amount
		case "up":
			aim -= amount
		case "down":
			aim += amount
		default:
			io.ReportError("", fmt.Errorf("unexpected command: %s", command))
		}
	}

	fmt.Println(x * y)
}
