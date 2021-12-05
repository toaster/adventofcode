package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/toaster/advent_of_code/internal/io"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	io.ReportError("failed to read standard input", err)

	x := 0
	y := 0
	aim := 0
	for _, line := range strings.Split(strings.Trim(string(input), "\n"), "\n") {
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
