package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	reportErr("failed to read standard input", err)

	var counts []int
	lineCount := 0
	for _, line := range strings.Split(strings.Trim(string(input), "\n"), "\n") {
		lineCount++
		for i, r := range line {
			var v int
			switch r {
			case '0':
				v = 0
			case '1':
				v = 1
			default:
				reportErr("", fmt.Errorf("unexpected bit: %c", r))
			}
			if i < len(counts) {
				counts[i] += v
			} else {
				counts = append(counts, v)
			}
		}
	}
	gamma := 0
	epsilon := 0
	for _, count := range counts {
		gamma = gamma << 1
		epsilon = epsilon << 1
		if count > lineCount/2 {
			gamma++
		} else {
			epsilon++
		}
	}

	fmt.Println(gamma * epsilon)
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
