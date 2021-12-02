package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	reportErr("failed to read standard input", err)

	x := 0
	y := 0
	for _, line := range strings.Split(strings.Trim(string(input), "\n"), "\n") {
		tokens := strings.Split(line, " ")
		if len(tokens) != 2 {
			reportErr("failed to parse input", fmt.Errorf("wrong token count: %d", len(tokens)))
		}

		command := tokens[0]
		amount, err := strconv.Atoi(tokens[1])
		reportErr("failed to parse input", err)

		switch command {
		case "forward":
			x += amount
		case "up":
			y -= amount
		case "down":
			y += amount
		default:
			reportErr("", fmt.Errorf("unexpected command: %s", command))
		}
	}

	fmt.Println(x * y)
}

func reportErr(msg string, err error) {
	if err != nil {
		if msg != "" {
			_, _ = fmt.Fprintln(os.Stderr, msg+":", err)
		} else {
			_, _ = fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
}
